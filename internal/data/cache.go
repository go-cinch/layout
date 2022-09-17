package data

import (
	"context"
	"fmt"
	"github.com/go-cinch/common/log"
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

// Cache .
type Cache struct {
	redis  redis.UniversalClient
	prefix string
	lock   string
	val    string
}

func (c *Cache) Register(prefix string) {
	c.prefix = prefix
	c.lock = fmt.Sprintf("%s_lock", prefix)
	c.val = fmt.Sprintf("%s_val", prefix)
}

func (c *Cache) Get(ctx context.Context, action string) (res string, ok bool) {
	var err error
	res, err = c.redis.Get(context.Background(), fmt.Sprintf("%s_%s", c.val, action)).Result()
	if err == nil {
		ok = true
	} else {
		log.
			WithContext(ctx).
			WithError(err).
			WithFields(log.Fields{
				"action": action,
			}).
			Warn("get cache failed")
	}
	return
}

func (c *Cache) Set(ctx context.Context, action, data string) {
	// set random expiration avoid a large number of keys expire at the same time
	seconds := rand.New(rand.NewSource(time.Now().Unix())).Int63n(300) + 300
	c.SetWithExpiration(ctx, action, data, seconds)
}

func (c *Cache) SetWithExpiration(ctx context.Context, action, data string, seconds int64) {
	// set random expiration avoid a large number of keys expire at the same time
	err := c.redis.Set(context.Background(), fmt.Sprintf("%s_%s", c.val, action), data, time.Duration(seconds)*time.Second).Err()
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithFields(log.Fields{
				"action":  action,
				"seconds": seconds,
			}).
			Warn("set cache failed")
	}
}

func (c *Cache) Flush(ctx context.Context, handler func(ctx context.Context) error) (err error) {
	err = handler(ctx)
	if err != nil {
		return
	}
	action := c.prefix + "*"
	arr := c.redis.Keys(context.Background(), action).Val()
	p := c.redis.Pipeline()
	for _, item := range arr {
		if item == c.lock {
			continue
		}
		p.Del(context.Background(), item)
	}
	_, pErr := p.Exec(context.Background())
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
	ok, _ = c.redis.SetNX(context.Background(), fmt.Sprintf("%s_%s", c.lock, action), 1, time.Minute).Result()
	retry := 0
	for retry < 100 && !ok {
		time.Sleep(10 * time.Millisecond)
		ok, _ = c.redis.SetNX(context.Background(), fmt.Sprintf("%s_%s", c.lock, action), 1, time.Minute).Result()
		retry++
	}
	return
}

func (c *Cache) Unlock(ctx context.Context, action string) {
	err := c.redis.Del(context.Background(), fmt.Sprintf("%s_%s", c.lock, action)).Err()
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

// NewCache .
func NewCache(client redis.UniversalClient) biz.Cache {
	return &Cache{
		redis: client,
	}
}
