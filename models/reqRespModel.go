package models

type AddTaskRequest struct {
	TaskData TaskModel `json:"task"`
}

type AddTaskResponse struct {
	InsertedId string `json:"insertedId"`
}

type GetListRequest struct {
	Page    int64  `json:"page" binding:"required"`
	PerPage int64  `json:"perPage" binding:"required"`
	Keyword string `json:"keyword"`
}

type GetListResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []TaskData `json:"data"`
}

type TaskData struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	CreatedEmail string    `json:"createdEmail"`
	CreatedName  string    `json:"createdName"`
	CreatedDated string    `json:"createdDated"`
	Comment      []Comment `json:"comments"`
}

type AddCommentRequest struct {
	Id      string  `json:"id"`
	Comment Comment `json:"comment"`
}

type SubmitStatusRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type CommonResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}
