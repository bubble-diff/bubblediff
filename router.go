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
	// RESTful API 开放接口
	// API接口仅在登录态有权限调用
	_apiv1 := r.Group("/api/v1")
	_apiv1.Use(JwtAuthMws.MiddlewareFunc())
	{
		// 任务详情 curd api
		_task := _apiv1.Group("/task")
		_task.GET("/:id", handlers.GetTaskDetailByID)
		_task.POST("", handlers.AddTask)
		_task.DELETE("/:id", handlers.DeleteTaskHandler.DeleteTask)
		_task.PUT("/:id", handlers.UpdateTaskHandler.UpdateTask)
		_task.GET("/searchByName", handlers.GetTaskDetailByName)

		// 开始diff任务 api
		_task.GET("/:id/start", handlers.StartTaskHandler.StartTask)
		_task.GET("/:id/stop", handlers.StopTaskHandler.StopTask)

		// 任务列表api
		_tasks := _apiv1.Group("/tasks")
		_tasks.GET("", handlers.GetTasks)

		// 查询当前请求用户的个人信息
		_user := _apiv1.Group("/userinfo")
		_user.GET("", handlers.GetUser)

		// Record记录 api
		_record := _apiv1.Group("/records")
		// 获取某个task的所有record meta信息
		_record.GET("/:taskid", handlers.ListRecordsMeta)
		// 获取特定记录详情
		_record.GET("/:taskid/:recordid", handlers.GetRecordDetail)

	}
	return r, nil
}

// conf := config.Get()
// store := sessions.NewCookieStore([]byte(conf.ClientSecret))
// r.Use(sessions.Sessions("mysession", store))
//
