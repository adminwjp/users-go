#####  reids cache example
```go
type RedisCache struct {

}
func(*RedisCache)Set(key string,val interface{})bool{
    return false;
}
```