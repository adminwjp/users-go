package tests

import (
	"bytes"
	"encoding/json"
	web_gin_controller "github.com/adminwjp/users-go/webs/gins"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"

	//"strconv"
	"sync"
	"testing"
)
var router *gin.Engine
//users\user_sample_test.go:16:9: undefined: RegisterGinRouter
func  init1()  {
	router = web_gin_controller.GinRouterImpl.RegisterGinRouter(nil,0)
}
// go test -count=1 -v  users/test/user_sample_test.go
//warn
//conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
func _TestRegister1(t *testing.T) {
	var w *httptest.ResponseRecorder

	assert := assert.New(t)

	// 1.测试 test 请求
	urlIndex := "/test"
	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
	return
	var wg sync.WaitGroup=sync.WaitGroup{} // 定义wg, 用来阻塞 goroutine
	var i int64 = 1000*1000*10000
	var m int64 = 1000*1000*10000*10

	for ; i < m; i++ {

		// 开一个等待
		wg.Add(1)

		go func(i int64) { // i 不属于临界资源，是安全的
			defer wg.Done() // 一个 goroutine 跑完后要减1，

			// 测试 /user/import 请求，模拟从 form 表单中获取数据
			param := make(map[string]string)
			p:="13"+strconv.FormatInt(i,10)
			param["phone"] =p
			param["pwd"] = p
			urlImport := "/register"
			if i%10==0{
				w = PostForm(urlImport, param, router)
				assert.Equal(200, w.Code)
				t.Log("register ==> "+p +w.Body.String())
			}else if i%5==0{
				param1 := make(map[string]interface{})
				param1["phone"] =p
				param1["pwd"] = p
				w = PostJson(urlImport, param1, router)
				assert.Equal(200, w.Code)
				t.Log("register ==> "+p +w.Body.String())
			}


		}(i)
	}
	// 等待上面的协程运行完，再接着测试
	wg.Wait()




	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
}

func _TestUpdate(t *testing.T) {
	var w *httptest.ResponseRecorder

	assert := assert.New(t)

	// 1.测试 test 请求
	urlIndex := "/test"
	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
	return
	var wg sync.WaitGroup=sync.WaitGroup{} // 定义wg, 用来阻塞 goroutine
	var i int64 = 1000*1000*10000
	var m int64 = 1000*1000*10000*10
	for ; i < m; i++ {

		// 开一个等待
		wg.Add(1)
		go func(i int64) { // i 不属于临界资源，是安全的
			defer wg.Done() // 一个 goroutine 跑完后要减1，

			p:="13"+strconv.FormatInt(i,10)
			urlImport := "/update/phone"
			w = Get(urlImport+"?account="+p+"&phone=13"+strconv.FormatInt(i+1,10), router)
			assert.Equal(200, w.Code)
			t.Log("update ==> "+p +w.Body.String())


		}(i)
	}
	// 等待上面的协程运行完，再接着测试
	wg.Wait()




	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
}

func _TestLogin(t *testing.T) {
	var w *httptest.ResponseRecorder

	assert := assert.New(t)

	// 1.测试 test 请求
	urlIndex := "/test"
	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
	return

	var wg sync.WaitGroup=sync.WaitGroup{} // 定义wg, 用来阻塞 goroutine
	var i int64 = 1000*1000*10000
	var m int64 = 1000*1000*10000*10
	for ; i < m; i++ {

		// 开一个等待
		wg.Add(1)
		go func(i int64) { // i 不属于临界资源，是安全的
			defer wg.Done() // 一个 goroutine 跑完后要减1，

			// 测试 /user/import 请求，模拟从 form 表单中获取数据
			param := make(map[string]string)
			p:="13"+strconv.FormatInt(i,10)
			param["phone"] =p
			param["pwd"] = p
			urlImport := "/login"
			if i%10==0{
				w = PostForm(urlImport, param, router)
				assert.Equal(200, w.Code)
				t.Log("login ==> "+p +w.Body.String())
			}else if i%5==0{
				param1 := make(map[string]interface{})
				param1["phone"] =p
				param1["pwd"] = p
				w = PostJson(urlImport, param1, router)
				assert.Equal(200, w.Code)
				t.Log("login ==> "+p +w.Body.String())
			}


		}(i)
	}
	// 等待上面的协程运行完，再接着测试
	wg.Wait()




	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)

}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

func Get(uri string, router1 *gin.Engine) *httptest.ResponseRecorder {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	return w
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	router.ServeHTTP(w, req)
	return w
}

// PostJson 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func PostJson(uri string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(param)
	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	return w
}
