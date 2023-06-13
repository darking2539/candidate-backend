package services

import (
	"candidate-backend/models"
	repo "candidate-backend/repositories"
	"errors"
	"time"
)

func AddTaskMockUpSv(taskModel models.TaskModel) (resp models.AddTaskResponse, err error) {

	taskModel.CreatedDated = time.Now()
	insertId, err := repo.CreateTaskRepo(taskModel)
	if err != nil {
		return
	}

	resp = models.AddTaskResponse{
		InsertedId: insertId,
	}

	return
}

func GetTaskListSv(page int64, perPage int64, keyword string) (resp models.GetListResponse, err error) {

	taskModel, pagination, err := repo.GetTaskListRepo(page, perPage, keyword)
	if err != nil {
		return
	}


	dataResp := []models.TaskData{}

	for _, task := range taskModel {
		dataResp = append(dataResp, models.TaskData{
			Id: task.ID.Hex(),
			Title: task.Title,
			Description: task.Description,
			Status: task.Status,
			CreatedEmail: task.CreatedEmail,
			CreatedName: task.CreatedName,
			CreatedDated: task.CreatedDated,
		})
	}

	resp = models.GetListResponse{
		Data: dataResp,
		Pagination: pagination,
	}

	return
}

func GetDetailSv(taskId string) (resp models.TaskModel, err error) {

	taskData, err := repo.GetTaskDetailRepo(taskId)
	if err != nil {
		return
	}

	resp = taskData

	return
}

func AddCommentSv(taskId string, comment models.Comment) (resp models.CommonResponse, err error) {

	comment.CreatedDated = time.Now()
	taskData, err := repo.AddCommentRepo(taskId, comment)
	if err != nil {
		return
	}

	if taskData.ID.Hex() == "000000000000000000000000" {
		err = errors.New("taskId not found")
		return
	}
	
	resp.Id = taskData.ID.Hex()
	resp.Status = "Add comment Sucessful"

	return
}

func KeepTaskSv(taskId string) (resp models.CommonResponse, err error) {

	taskData, err := repo.KeepTaskRepo(taskId)
	if err != nil {
		return
	}

	if taskData.ID.Hex() == "000000000000000000000000" {
		err = errors.New("taskId not found")
		return
	}
	
	resp.Id = taskData.ID.Hex()
	resp.Status = "Keep Task Sucessful"

	return
}

func SubmitStatusSv(taskId string, status string) (resp models.CommonResponse, err error) {

	taskData, err := repo.SubmitStatusRepo(taskId, status)
	if err != nil {
		return
	}

	if taskData.ID.Hex() == "000000000000000000000000" {
		err = errors.New("taskId not found")
		return
	}
	
	resp.Id = taskData.ID.Hex()
	resp.Status = "Submit Status Sucessful"

	return
}