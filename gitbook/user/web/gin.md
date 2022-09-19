##### gin example
```go
func (*GinRouter)RegisterGinRouter(gin1 *gin.Engine ,port int)*gin.Engine  {
	var r *gin.Engine
	if gin1!=nil{
		r=gin1
	}else{
		r= gin.Default()
		r.Use(gins.RegisterCors())
		r.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": false,"code":404,"msg":"resource not found "})
			c.Abort()
		})

		r.NoMethod(func(c *gin.Context) {
			defer func() {
				if r:=recover();r!=nil{
					c.JSON(http.StatusOK, gin.H{"status": false,"code":500,"msg":"server inner error "})
				}

			}()
			c.Next()
		})
		r.GET("/test", func(c *gin.Context){
			c.JSON(200,gin.H{
				"status": true, "code": 200, "msg": "success",
			})
		})
	}
}
```