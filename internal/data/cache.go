package data

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/plugins/gorm/tenant"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Cache .
type Cache struct {
	redis   redis.UniversalClient
	locker  *redislock.Client
	disable bool
	prefix  string
	lock    string
	val     string
	refresh bool
}

var local = cache.New(30*time.Minute, 60*time.Minute)

// NewCache .
func NewCache(c *conf.Bootstrap, client redis.UniversalClient) biz.Cache {
	return &Cache{
		redis:   client,
		locker:  redislock.New(client),
		disable: c.Server.Nocache,
		lock:    "lock",
		val:     "val",
	}
}

func (c *Cache) Cache() redis.UniversalClient {
	return c.redis
}

func (c *Cache) WithPrefix(prefix string) biz.Cache {
	return &Cache{
		redis:   c.redis,
		locker:  c.locker,
		disable: c.disable,
		prefix:  prefix,
		lock:    c.lock,
		val:     c.val,
	}
}

func (c *Cache) WithRefresh() biz.Cache {
	return &Cache{
		redis:   c.redis,
		locker:  c.locker,
		disable: c.disable,
		prefix:  c.prefix,
		lock:    c.lock,
		val:     c.val,
		refresh: true,
	}
}

func (c *Cache) Get(
	ctx context.Context,
	action string,
	write func(context.Context) (string, error),
) (res string, err error) {
	tr := otel.Tracer("cache")
	ctx, span := tr.Start(ctx, "Get")
	defer span.End()
	if c.disable {
		return write(ctx)
	}
	key := c.getValKey(ctx, action)
	if !c.refresh {
		// 1. first get cache
		// 1.1. get from local
		res, err = GetFromLocal(key)
		if err == nil {
			// cache exists
			return
		}
		// 1.2. get from redis
		res, err = c.redis.Get(ctx, key).Result()
		if err == nil {
			// cache exists
			return
		}
	}
	// 2. get lock before read db
	lock, err := c.Lock(ctx, action)
	if err != nil {
		err = biz.ErrTooManyRequests(ctx)
		return
	}
	defer func() {
		_ = lock.Release(ctx)
	}()
	if !c.refresh {
		// 3. double check cache exists(avoid concurrency step 1 ok=false)
		res, err = GetFromLocal(key)
		if err == nil {
			// cache exists
			return
		}
		res, err = c.redis.Get(ctx, key).Result()
		if err == nil {
			// cache exists
			return
		}
	}
	// 4. load data from db and write to cache
	if write != nil {
		res, err = write(ctx)
	}
	return
}

func (c *Cache) Set(ctx context.Context, action, data string, short bool) {
	// set random expiration avoid a large number of keys expire at the same time
	seconds := rand.New(rand.NewSource(time.Now().Unix())).Int63n(300) + 300
	if short {
		// if record not found, set a short expiration
		seconds = 60
	}
	c.SetWithExpiration(ctx, action, data, seconds)
}

func (c *Cache) SetWithExpiration(ctx context.Context, action, data string, seconds int64) {
	if c.disable {
		return
	}
	key := c.getValKey(ctx, action)
	// set to local cache
	Set2Local(key, data, int(seconds))
	// set to redis
	err := c.redis.Set(ctx, key, data, time.Duration(seconds)*time.Second).Err()
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithFields(log.Fields{
				"action":  action,
				"seconds": seconds,
			}).
			Warn("set cache failed")
		return
	}
}

func (c *Cache) Del(ctx context.Context, action string) {
	if c.disable {
		return
	}
	key := c.getValKey(ctx, action)
	DelFromLocal(key)
	err := c.redis.Del(ctx, key).Err()
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithFields(log.Fields{
				"action": action,
				"key":    key,
			}).
			Warn("del cache failed")
	}
}

func (c *Cache) Flush(ctx context.Context, handler func(ctx context.Context) error) (err error) {
	err = handler(ctx)
	if err != nil {
		return
	}
	if c.disable {
		return
	}
	action := c.getPrefixKey(ctx)
	arr := c.redis.Keys(ctx, action).Val()
	p := c.redis.Pipeline()
	for _, item := range arr {
		if item == c.lock {
			continue
		}
		DelFromLocal(item)
		p.Del(ctx, item)
	}
	_, pErr := p.Exec(ctx)
	if pErr != nil {
		log.
			WithContext(ctx).
			WithError(pErr).
			WithFields(log.Fields{
				"action": action,
			}).
			Warn("flush cache failed")
	}
	return
}

func (c *Cache) FlushByPrefix(ctx context.Context, prefix ...string) (err error) {
	action := c.getPrefixKey(ctx, prefix...)
	arr := c.redis.Keys(ctx, action).Val()
	p := c.redis.Pipeline()
	for _, item := range arr {
		if item == c.lock {
			continue
		}
		DelFromLocal(item)
		p.Del(ctx, item)
	}
	_, pErr := p.Exec(ctx)
	if pErr != nil {
		log.
			WithContext(ctx).
			WithError(pErr).
			WithFields(log.Fields{
				"action": action,
			}).
			Warn("flush cache by prefix failed")
	}
	return
}

func (c *Cache) Lock(ctx context.Context, action string) (*redislock.Lock, error) {
	tr := otel.Tracer("cache")
	ctx, span := tr.Start(ctx, "Lock")
	defer span.End()
	lock, err := c.locker.Obtain(
		ctx,
		c.getLockKey(ctx, action),
		20*time.Second,
		&redislock.Options{
			RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(5*time.Millisecond), 400),
		},
	)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}
	return lock, err
}

func Set2Local(key, val string, expire int) {
	local.Set(key, val, time.Duration(expire)*time.Second)
}

func GetFromLocal(key string) (string, error) {
	val, ok := local.Get(key)
	if !ok {
		return "", errors.New("key not found")
	}
	return val.(string), nil
}

func DelFromLocal(key string) {
	local.Delete(key)
}

func (c *Cache) getPrefixKey(ctx context.Context, arr ...string) string {
	id := tenant.FromContext(ctx)
	prefix := c.prefix
	if len(arr) > 0 {
		// append params prefix need add val
		prefix = strings.Join(append([]string{prefix, c.val}, arr...), "_")
	}
	if strings.TrimSpace(prefix) == "" {
		// avoid flush all key
		log.
			WithContext(ctx).
			Warn("invalid prefix")
		prefix = "prefix"
	}
	if id == "" {
		return strings.Join([]string{prefix, "*"}, "")
	}
	return strings.Join([]string{id, "_", prefix, "*"}, "")
}

func (c *Cache) getValKey(ctx context.Context, action string) string {
	id := tenant.FromContext(ctx)
	if id == "" {
		return strings.Join([]string{c.prefix, c.val, action}, "_")
	}
	return strings.Join([]string{id, c.prefix, c.val, action}, "_")
}

func (c *Cache) getLockKey(ctx context.Context, action string) string {
	id := tenant.FromContext(ctx)
	if id == "" {
		return strings.Join([]string{c.prefix, c.lock, action}, "_")
	}
	return strings.Join([]string{id, c.prefix, c.lock, action}, "_")
}
