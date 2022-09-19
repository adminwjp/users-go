package cache_redis_impl

import (
	"encoding/json"
	cache_redis "github.com/adminwjp/infrastructure-go/caches/redises"
	"github.com/adminwjp/users-go/models"
	"github.com/go-redis/redis"
	"log"
)

type RedisCache struct {
	cache_redis.RedisCache
	roles []*models.RoleModel
	emails []*models.EmailConfigModel
	smses []*models.SmsConfigModel
	paies []*models.PaySecrtConfigModel
}

func(cache *RedisCache)  Init()  {
	sub:=cache.Client.Subscribe("roles","emails","smses","paies")
	iface,err:=sub.Receive()
	if err!=nil{
		log.Printf("subscribe fail,err:%s",err.Error())
		return
	}
	//    // Should be *Subscription, but others are possible if other actions have been
	//    // taken on sub since it was created.
	    switch iface.(type) {
			case *redis.Subscription:
				// subscribe succeeded
				log.Println("subscribe succeeded")
				s:=iface.(*redis.Subscription)
				if s!=nil{

				}
				break
			case *redis.Message:
				// received first message
				log.Println("received first message")
				m:=iface.(*redis.Message)
				if m!=nil{
					switch m.Channel {
					case "roles":
						json.Unmarshal([]byte(m.String()),&cache.roles)
						if cache.roles==nil{
							cache.roles=make([]*models.RoleModel,0)
						}
						break
					case "emails":json.Unmarshal([]byte(m.String()),&cache.emails)
						if cache.emails==nil{
							cache.emails=make([]*models.EmailConfigModel,0)
						}
						break
					case "smses":json.Unmarshal([]byte(m.String()),&cache.smses)
						if cache.smses==nil{
							cache.smses=make([]*models.SmsConfigModel,0)
						}
						break
					case "paies":json.Unmarshal([]byte(m.String()),&cache.paies)
						if cache.paies==nil{
							cache.paies=make([]*models.PaySecrtConfigModel,0)
						}
						break
					default:
						break
					}
				}
				break
			case *redis.Pong:
				// pong received
				log.Println("pong received")
				p:=iface.(*redis.Pong)
				if p!=nil{

				}
				break
		   default:
			   // handle error
			   log.Println("handle error")
			   break
	    }

	    //ch := sub.Channel()
	    //m1:=&redis.Message{}
	    //ch <- m1
}
func(cache *RedisCache) SaveRole(m  *models.RoleModel)  {
	e :=false
	for i := 0; i < len(cache.roles); i++ {
		if cache.roles[i].Id==m.Id{
			e=true
			cache.roles[i]=m
			break
		}
	}
	if !e{
		cache.roles=append(cache.roles,m)
	}
	bu,_:=json.Marshal(cache.roles)
	cache.Client.Publish("roles",string(bu))
}
func(cache *RedisCache) SaveRoles(ms  []*models.RoleModel)  {
	bu,_:=json.Marshal(ms)
	cache.Client.Publish("roles",string(bu))
}
func(cache *RedisCache) SaveEmail(m  *models.EmailConfigModel)  {
	e :=false
	for i := 0; i < len(cache.emails); i++ {
		if cache.emails[i].Id==m.Id{
			e=true
			cache.emails[i]=m
			break
		}
	}
	if !e{
		cache.emails=append(cache.emails,m)
	}
	bu,_:=json.Marshal(cache.emails)
	cache.Client.Publish("emails",string(bu))
}
func(cache *RedisCache) SaveSmses(ms  []*models.SmsConfigModel)  {
	bu,_:=json.Marshal(ms)
	cache.Client.Publish("smses",string(bu))
}
func(cache *RedisCache) SaveSms(m  *models.SmsConfigModel)  {
	e :=false
	for i := 0; i < len(cache.smses); i++ {
		if cache.smses[i].Id==m.Id{
			e=true
			cache.smses[i]=m
			break
		}
	}
	if !e{
		cache.smses=append(cache.smses,m)
	}
	bu,_:=json.Marshal(cache.smses)
	cache.Client.Publish("smses",string(bu))
}
func(cache *RedisCache) SavePaies(ms  []*models.PaySecrtConfigModel)  {
	bu,_:=json.Marshal(ms)
	cache.Client.Publish("paies",string(bu))
}
func(cache *RedisCache) SavePay(m  *models.PaySecrtConfigModel)  {
	e :=false
	for i := 0; i < len(cache.paies); i++ {
		if cache.paies[i].Id==m.Id{
			e=true
			cache.paies[i]=m
			break
		}
	}
	if !e{
		cache.paies=append(cache.paies,m)
	}
	bu,_:=json.Marshal(cache.paies)
	cache.Client.Publish("paies",string(bu))
}