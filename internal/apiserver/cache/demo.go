package cache

import (
	"context"
	"github.com/ashwinyue/maltx/pkg/distlock"
)

type DemoCache interface {
	Create(ctx context.Context, key string) error
}

type demoCache struct {
	datacache *datacache
	keyPrefix string
}

var _ DemoCache = (*demoCache)(nil)

// newUserStore 创建 userStore 的实例.
func newDemoCache(cache *datacache) *demoCache {
	return &demoCache{
		datacache: cache,
		keyPrefix: "demoKey",
	}
}

func (s *demoCache) Create(ctx context.Context, key string) error {
	//if err := s.store.DB(ctx).Create(&obj).Error; err != nil {
	//	log.Errorw("Failed to insert post into database", "err", err, "post", obj)
	//	return errno.ErrDBWrite.WithMessage(err.Error())
	//}
	s.datacache.rds.Del(ctx, key)

	return nil
}

func (s *demoCache) GetLocker() *distlock.RedisLocker {
	return distlock.NewRedisLocker(s.datacache.RDS(), distlock.WithLockName(s.keyPrefix))
}
