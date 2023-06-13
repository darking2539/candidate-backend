package router

import (
	ctl "candidate-backend/controller"
	"candidate-backend/db"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	router.GET("/healthz", db.Healthz)

	//mockup
	router.POST("/task/add", ctl.AddTaskMockUpCtl)

	//viewer
	router.GET("/task/:id", ctl.GetTaskDetailCtl)
	router.POST("/task/list", ctl.GetTaskListCtl)
	router.POST("/task/comment", ctl.AddCommentToTaskCtl)
	router.POST("/task/submit", ctl.SubmitStatusTaskCtl)
	router.POST("/task/keep/:id", ctl.KeepTaskCtl)

}
