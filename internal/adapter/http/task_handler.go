package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type TaskHandlers struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandlers {
	return &TaskHandlers{
		Service: service,
	}
}

func (h *TaskHandlers) CreateTask(c *gin.Context) {

	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	var body request.Task
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	task := domain.Task{
		Title:       body.Title,
		Description: body.Description,
		Category_Id: body.Category_Id,
	}

	result, err := h.Service.CreateTask(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreateTask{
		ID:          result.ID,
		Title:       result.Title,
		Status:      result.Status,
		Description: result.Description,
		User_Id:     result.User_Id,
		Category_Id: result.Category_Id,
		Created_At:  result.Created_At,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TaskHandlers) GetTask(c *gin.Context) {

	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	result, err := h.Service.GetTask(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var tasks response.Tasks

	for _, task := range *result {
		tasks.Tasks = append(tasks.Tasks, &response.Task{
			ID:          task.ID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			User_Id:     task.User_Id,
			Category_Id: task.Category_Id,
			Created_At:  task.Created_At,
			User: response.Users{
				ID:        task.Users.ID,
				Email:     task.Users.Email,
				Full_Name: task.Users.Full_name,
			},
		})
	}

	c.JSON(http.StatusOK, tasks.Tasks)
}

func (h *TaskHandlers) UpdateTask(c *gin.Context) {
	var body request.Task
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	urlParam := c.Param("id")
	taskID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	task := domain.Task{
		ID:          uint(taskID),
		Title:       body.Title,
		Description: body.Description,
	}

	result, err := h.Service.UpdateTask(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateTask{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Status:      result.Status,
		User_Id:     result.User_Id,
		Category_Id: result.Category_Id,
		Updated_At:  result.Updated_At,
	})
}

func (h *TaskHandlers) UpdateStatusTask(c *gin.Context) {
	var body request.Task
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	urlParam := c.Param("id")
	taskID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	task := domain.Task{
		ID:     uint(taskID),
		Status: body.Status,
	}

	result, err := h.Service.UpdateStatusTask(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateStatusTask{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Status:      result.Status,
		User_Id:     result.User_Id,
		Category_Id: result.Category_Id,
		Updated_At:  result.Updated_At,
	})
}

func (h *TaskHandlers) UpdateCategoryTask(c *gin.Context) {

}

func (h *TaskHandlers) DeleteTask(c *gin.Context) {

}
