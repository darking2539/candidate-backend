package controller

import (
	"candidate-backend/models"
	sv "candidate-backend/services"
	"encoding/json"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)


func AddTaskMockUpCtl(c *gin.Context) {

	var request models.AddTaskRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}
	
	svResp, svErr := sv.AddTaskMockUpSv(request.TaskData)
	if svErr != nil {
		boom.Internal(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func GetTaskListCtl(c *gin.Context) {

	var request models.GetListRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}
	
	svResp, svErr := sv.GetTaskListSv(request.Page, request.PerPage, request.Keyword)
	if svErr != nil {
		boom.Internal(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func GetTaskDetailCtl(c *gin.Context) {

	taskId := c.Param("id") 
	
	svResp, svErr := sv.GetDetailSv(taskId)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func AddCommentToTaskCtl(c *gin.Context) {
	
	var request models.AddCommentRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}
	
	svResp, svErr := sv.AddCommentSv(request.Id, request.Comment)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func KeepTaskCtl(c *gin.Context) {
	
	taskId := c.Param("id") 
	
	svResp, svErr := sv.KeepTaskSv(taskId)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func SubmitStatusTaskCtl(c *gin.Context) {
	
	var request models.SubmitStatusRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}
	
	svResp, svErr := sv.SubmitStatusSv(request.Id, request.Status)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	resp := svResp
	
	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}