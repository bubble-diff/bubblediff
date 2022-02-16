package main

import (
	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/handlers"
)

// InitRouter 注册路由
func InitRouter() (r *gin.Engine, err error) {
	err = InitMiddlewares()
	if err != nil {
		return nil, err
	}
	r = gin.Default()
	r.Use(CorsMws)

	// 测试接口
	r.GET("/ping", handlers.Ping)
	// 登录相关
	r.GET("/login", JwtAuthMws.LoginHandler)
	// API 开放接口
	// API接口仅在登录态有权限调用
	_apiv1 := r.Group("/api/v1")
	_apiv1.Use(JwtAuthMws.MiddlewareFunc())
	{
		// Diff任务curd
		_task := _apiv1.Group("/task")
		_task.POST("/add")
		_task.POST("/delete")
		_task.GET("/get")
		_task.POST("/update")
		// User curd
		_user := _apiv1.Group("/user")
		_user.GET("/get", handlers.GetUser)
	}
	return r, nil
}

// conf := config.Get()
// store := sessions.NewCookieStore([]byte(conf.ClientSecret))
// r.Use(sessions.Sessions("mysession", store))
//
