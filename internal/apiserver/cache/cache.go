package cache

import (
	redisstore "github.com/ashwinyue/maltx/pkg/cache/store/redis"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"sync"
)

var ProviderSet = wire.NewSet(NewDataCache, wire.Bind(new(ICache), new(*datacache)))

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 datastore 实例.
	S *datacache
)

type ICache interface {
	// 返回 Store 层的 *gorm.DB 实例，在少数场景下会被用到.
	//RDS(ctx context.Context) *redisstore.RedisStore

	Demo() DemoCache
}

type datacache struct {
	rds *redisstore.RedisStore
}

func (datacache *datacache) RDS() *redis.Client {
	return datacache.rds.RDS()
}

func NewDataCache(rds *redisstore.RedisStore) *datacache {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datacache{rds: rds}
	})

	return S
}

// Users 返回一个实现了 UserStore 接口的实例.
func (cache *datacache) Demo() DemoCache {
	return newDemoCache(cache)
}
