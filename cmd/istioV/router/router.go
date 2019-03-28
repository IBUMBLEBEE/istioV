package router

import (
	"IBUMBLEBEE/istioV/cmd/istioV/handler"

	"github.com/gin-gonic/gin"
)

// Executor init route func
func Executor(engine *gin.Engine) {
	engine.Use(
		gin.Recovery(),
	)

	// public APIs
	engine.GET("/api/proc/healthy", handler.Healthy)
	// for login interface
	// engine.POST("/api/login/accunt")
	api := engine.Group("/api")

	// for identification interface
	// api.Use(middleware.Auth())
	api.GET("/dashboard", handler.DisplayDashboard)

	// api.GET("/services", handler.Service)
	// api.GET("/services/:uid", handler.Service)
	// api.GET("/services/:uid/pod", handler.Service)

	// api.GET("/pods", handler.Pods)
	// api.GET("/pods/:name", handler.Pods)

	// api.GET("/tasks", handler.ListTasks)
	// api.GET("/tasks/:id", handler.ListTasks)
	// api.POST("/tasks", handler.AddTasks)

	// monitos interface
	// api.GET("/metrics", handler.DisplayMetrics)
	// api.GET("/traffic", handler.Traffic)
	// api.GET("/d3graph", handler.D3Graph)

	// Istio config
	// api.GET("/tasktmpls", handler.ListTaskTemplates)
	// api.GET("/tasktmpls/:id", handler.ListTaskTemplates)
	// api.POST("/addtasktmpls", handler.AddTaskTaskTemplates)
	// api.PUT("/tasktmpls/:id", handler.UpdateTaskTemplates)
	// api.DELETE("/tasktmpls/:id", handler.DeleteTaskTemplates)
	// api.GET("/tasktmpls/:id/vars", handler.ListTaskTemplatesVars)
	// api.GET("/kube/info", handler.KubeInfo)
}
