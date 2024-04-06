package data

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/go-cinch/common/log"
	"github.com/go-cinch/common/plugins/gorm/tenant"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
)

// Cache .
type Cache struct {
	redis   redis.UniversalClient
	disable bool
	prefix  string
	lock    string
	val     string
}

// NewCache .
func NewCache(c *conf.Bootstrap, client redis.UniversalClient) biz.Cache {
	return &Cache{
		redis:   client,
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
		disable: c.disable,
		prefix:  prefix,
		lock:    c.lock,
		val:     c.val,
	}
}

func (c *Cache) Get(
	ctx context.Context,
	action string,
	write func(context.Context) (string, error),
) (res string, err error) {
	if c.disable {
		return write(ctx)
	}
	key := c.getValKey(ctx, action)
	// 1. first get cache
	res, err = c.redis.Get(ctx, key).Result()
	if err == nil {
		// cache exists
		return
	}
	// 2. get lock before read db
	ok := c.Lock(ctx, action)
	if !ok {
		err = biz.ErrTooManyRequests(ctx)
		return
	}
	defer c.Unlock(ctx, action)
	// 3. double check cache exists(avoid concurrency step 1 ok=false)
	res, err = c.redis.Get(ctx, key).Result()
	if err == nil {
		// cache exists
		return
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
	// set random expiration avoid a large number of keys expire at the same time
	err := c.redis.Set(ctx, c.getValKey(ctx, action), data, time.Duration(seconds)*time.Second).Err()
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

func (c *Cache) Lock(ctx context.Context, action string) (ok bool) {
	if c.disable {
		ok = true
		return
	}
	retry := 0
	var e error
	for retry < 20 && !ok {
		ok, e = c.redis.SetNX(ctx, c.getLockKey(ctx, action), 1, time.Minute).Result()
		if errors.Is(e, context.DeadlineExceeded) ||
			errors.Is(e, context.Canceled) ||
			(e != nil && e.Error() == "redis: connection pool timeout") {
			log.
				WithContext(ctx).
				WithError(e).
				WithFields(log.Fields{
					"action": action,
				}).
				Warn("lock failed")
			return
		}
		time.Sleep(25 * time.Millisecond)
		retry++
	}
	return
}

func (c *Cache) Unlock(ctx context.Context, action string) {
	if c.disable {
		return
	}
	// get span and create new ctx since current ctx maybe timeout, unlock must be execution
	span := trace.SpanFromContext(ctx)
	ctx = trace.ContextWithSpan(context.Background(), span)
	err := c.redis.Del(ctx, c.getLockKey(ctx, action)).Err()
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithFields(log.Fields{
				"action": action,
			}).
			Warn("unlock cache failed")
	}
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
		return strings.Join([]string{prefix, "*"}, "_")
	}
	return strings.Join([]string{id, prefix, "*"}, "_")
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
