package hystrix

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Wrapper(
	size,
	reqThreshold int,
	failedThreshold float64,
	duration time.Duration,
) gin.HandlerFunc {
	r := NewRollingWindow(size, reqThreshold, failedThreshold, duration)
	// 启动了一个goroutine 来每隔200ms 增加一个 bucket,超过窗口大小就删除
	r.Launch()
	// 启动一个goroutine 来监控是否开启熔断
	r.Monitor()
	// 启动一个goroutine 每隔一秒打印熔断状态
	r.ShowStatus()
	// 变量可以在匿名函数中使用吗？查看变量的作用范围
	return func(c *gin.Context) {
		if r.Broken() {
			c.String(http.StatusInternalServerError, "请求被上游服务拒绝")
			c.Abort()
			return
		}
		c.Next()
		if c.Writer.Status() != http.StatusOK {
			r.RecordReqResult(false)
		} else {
			r.RecordReqResult(true)
		}
	}
}
