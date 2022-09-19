package cache_redis_impl

import "github.com/go-redis/redis"
var AdminCacheInstance=&AdminCacheImpl{}
var UserCacheInstance=&UserCacheImpl{}
func intResult(val *redis.IntCmd) (int,error) {
	if val==nil{
		return 0,nil
	}
	return int(val.Val()),val.Err()
}
