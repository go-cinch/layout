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
)

// Cache .
type Cache struct {
	redis  redis.UniversalClient
	prefix string
	lock   string
	val    string
}

// NewCache .
func NewCache(c *conf.Bootstrap, client redis.UniversalClient) biz.Cache {
	return &Cache{
		redis:  client,
		prefix: c.Name,
		lock:   "lock",
		val:    "val",
	}
}

func (c *Cache) Cache() redis.UniversalClient {
	return c.redis
}

func (c *Cache) WithPrefix(prefix string) biz.Cache {
	return &Cache{
		redis:  c.redis,
		prefix: strings.Join([]string{c.prefix, prefix}, "_"),
		lock:   c.lock,
		val:    c.val,
	}
}

func (c *Cache) Get(
	ctx context.Context,
	action string,
	write func(context.Context) (string, bool),
) (res string, ok bool) {
	ctx = getDefaultTimeoutCtx(ctx)
	key := c.getValKey(ctx, action)
	var err error
	// 1. first get cache
	res, err = c.redis.Get(ctx, key).Result()
	if err == nil {
		// cache exists
		ok = true
		return
	}
	// 2. get lock before read db
	ok = c.Lock(ctx, action)
	if !ok {
		return
	}
	defer c.Unlock(ctx, action)
	// 3. double check cache exists(avoid concurrency step 1 ok=false)
	res, err = c.redis.Get(ctx, key).Result()
	if err == nil {
		// cache exists
		ok = true
		return
	}
	// 4. load data from db and write to cache
	if write != nil {
		res, ok = write(ctx)
	}
	return
}

func (c *Cache) Set(ctx context.Context, action, data string, short bool) {
	ctx = getDefaultTimeoutCtx(ctx)
	// set random expiration avoid a large number of keys expire at the same time
	seconds := rand.New(rand.NewSource(time.Now().Unix())).Int63n(300) + 300
	if short {
		// if record not found, set a short expiration
		seconds = 60
	}
	c.SetWithExpiration(ctx, action, data, seconds)
}

func (c *Cache) SetWithExpiration(ctx context.Context, action, data string, seconds int64) {
	ctx = getDefaultTimeoutCtx(ctx)
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
	ctx = getDefaultTimeoutCtx(ctx)
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
	ctx = getDefaultTimeoutCtx(ctx)
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

func (c *Cache) Lock(ctx context.Context, action string) (ok bool) {
	retry := 0
	var e error
	for retry < 20 && !ok {
		ctx = getDefaultTimeoutCtx(ctx)
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
	ctx = getDefaultTimeoutCtx(ctx)
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

func (c *Cache) getPrefixKey(ctx context.Context) string {
	id := tenant.FromContext(ctx)
	prefix := c.prefix
	if strings.TrimSpace(c.prefix) == "" {
		// avoid flush all key
		log.
			WithContext(ctx).
			Warn("invalid prefix")
		prefix = "prefix"
	}
	return strings.Join([]string{id, prefix, "*"}, "_")
}

func (c *Cache) getValKey(ctx context.Context, action string) string {
	id := tenant.FromContext(ctx)
	return strings.Join([]string{id, c.prefix, c.val, action}, "_")
}

func (c *Cache) getLockKey(ctx context.Context, action string) string {
	id := tenant.FromContext(ctx)
	return strings.Join([]string{id, c.prefix, c.lock, action}, "_")
}
