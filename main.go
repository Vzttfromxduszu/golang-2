package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping/api"
	"shopping/utils/graceful"

	_ "shopping/docs"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 电商项目
// @description 电商项目
// @version 1.0
// @contact.name golang技术栈
// @contact.url https://www.golang-tech-stack.com

// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()
	registerMiddlewares(r)
	api.RegisterHandlers(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	graceful.ShutdownGin(srv, time.Second*5)

}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	//r.Use(...)：将一个或多个中间件添加到路由引擎中，使其在处理每个请求时都会执行
	r.Use(
		// 自定义日志格式的 Logger 中间件
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {

				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339),
					param.Method,
					param.Path,
					param.Request.Proto, // like "HTTP/1.0"
					param.StatusCode,
					param.Latency,
					param.ErrorMessage,
				)
			}))
	r.Use(gin.Recovery())
}

// type LogFormatterParams struct {
// 	Request *http.Request

// 	// TimeStamp shows the time after the server returns a response.
// 	TimeStamp time.Time
// 	// StatusCode is HTTP response code.
// 	StatusCode int
// 	// Latency is how much time the server cost to process a certain request.
// 	Latency time.Duration
// 	// ClientIP equals Context's ClientIP method.
// 	ClientIP string
// 	// Method is the HTTP method given to the request.
// 	Method string
// 	// Path is a path the client requests.
// 	Path string
// 	// ErrorMessage is set if error has occurred in processing the request.
// 	ErrorMessage string
// 	// isTerm shows whether gin's output descriptor refers to a terminal.
// 	isTerm bool
// 	// BodySize is the size of the Response Body
// 	BodySize int
// 	// Keys are the keys set on the request's context.
// 	Keys map[string]any
// }
