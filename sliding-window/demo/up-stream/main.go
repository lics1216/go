package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sliding-window/pkg/hystrix"
	"time"
)

func NewUpStreamServer(
	size,
	reqThreshold int,
	failedThreshold float64,
	duration time.Duration,
) *gin.Engine {
	app := gin.Default()
	// 后面两个参数是两个handler 函数
	app.GET("/api/up/v1", hystrix.Wrapper(
		size,
		reqThreshold,
		failedThreshold,
		duration,
	), upHandler)
	return app
}

func upHandler(c *gin.Context) {
	// 如果上游接到请求就会给下游发请求，
	res, err := http.Get("http://localhost:8000/api/down/v1")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		c.String(res.StatusCode, string(data))
		return
	}
	c.String(res.StatusCode, string(data))
}


func main() {
	// @todo 把配置提取到一个配置文件中，这样方便调参
	NewUpStreamServer(
		10,
		2,
		0.4,
		time.Second*5,
	).Run(":8001")
}
