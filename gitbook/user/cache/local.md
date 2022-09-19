##### cache example
```go
type LocalCache struct {

}
func(*LocalCache)Set(key string,val interface{})bool{
    return false;
}
```