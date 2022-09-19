package service_impl

import (
	"github.com/adminwjp/infrastructure-go/caches/redises"
	cache_impl "github.com/adminwjp/users-go/caches/impls"
	cache_redis_impl "github.com/adminwjp/users-go/caches/impls/redis"
	"github.com/adminwjp/users-go/datas"
)

type CacheConfigImpl struct {

}
var RedisCacheInstance *redises.RedisCache
func (config *CacheConfigImpl)CacheByAdmin(service1 *AdminServiceImpl)  {
	switch datas.GlobalConfig.CacheFlag {
	case datas.CacheEmpty,datas.CacheFileMap,datas.CacheLocal:
		service1.BaseCache=cache_impl.AdminCacheInstance
		service1.Cache=cache_impl.AdminCacheInstance
		break
	case datas.CacheRedis:
		service1.BaseCache=cache_redis_impl.AdminCacheInstance
		service1.Cache=cache_redis_impl.AdminCacheInstance
		cache_redis_impl.AdminCacheInstance.RedisCache=*RedisCacheInstance
		break
	case datas.CacheBee:
		service1.BaseCache=cache_impl.AdminCacheInstance
		service1.Cache=cache_impl.AdminCacheInstance
		break
	default:
		service1.BaseCache=cache_impl.AdminCacheInstance
		service1.Cache=cache_impl.AdminCacheInstance

		break
	}

}

func (config *CacheConfigImpl)CacheByUser(service1 *UserServiceImpl)  {
	switch datas.GlobalConfig.CacheFlag {
	case datas.CacheEmpty,datas.CacheFileMap,datas.CacheLocal:
		service1.BaseCache=cache_impl.UserCacheInstance
		service1.Cache=cache_impl.UserCacheInstance
		break
	case datas.CacheRedis:
		service1.BaseCache=cache_redis_impl.UserCacheInstance
		service1.Cache=cache_redis_impl.UserCacheInstance
		cache_redis_impl.UserCacheInstance.RedisCache=*RedisCacheInstance
		break
	case datas.CacheBee:
		service1.BaseCache=cache_impl.UserCacheInstance
		service1.Cache=cache_impl.UserCacheInstance
		break
	default:
		service1.BaseCache=cache_impl.UserCacheInstance
		service1.Cache=cache_impl.UserCacheInstance
		break
	}

}
