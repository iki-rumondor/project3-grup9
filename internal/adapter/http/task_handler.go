package customHTTP

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project3-grup9/internal/application"
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {

	var body request.CreateTask
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "your request body is not valid",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")

	task := domain.Task{
		Title:       body.Title,
		Description: body.Description,
		CategoryID:  body.CategoryID,
		UserID:      userID,
		Status:      false,
	}

	result, err := h.Service.CreateTask(&task)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

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
		UserID:      result.UserID,
		CategoryID:  result.CategoryID,
		CreatedAt:   result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {

	result, err := h.Service.GetTasks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var tasks = []*response.Task{}

	for _, task := range *result {
		tasks = append(tasks, &response.Task{
			ID:          task.ID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserID:      task.UserID,
			CategoryID:  task.CategoryID,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			User: response.UserTask{
				ID:       task.User.ID,
				Email:    task.User.Email,
				FullName: task.User.FullName,
			},
		})
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var body request.UpdateTask
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "your request body is not valid",
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

	userID := c.GetUint("user_id")

	task := domain.Task{
		ID:          uint(taskID),
		Title:       body.Title,
		Description: body.Description,
		UserID:      userID,
	}

	result, err := h.Service.UpdateTask(&task)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

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
		UserID:      result.UserID,
		CategoryID:  result.CategoryID,
		UpdatedAt:   result.UpdatedAt,
	})
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	var body request.UpdateTaskStatus
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "your request body is not valid",
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

	userID := c.GetUint("user_id")

	task := domain.Task{
		ID:     uint(taskID),
		Status: body.Status,
		UserID: userID,
	}

	result, err := h.Service.UpdateTaskStatus(&task)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

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
		UserID:      result.UserID,
		CategoryID:  result.CategoryID,
		UpdatedAt:   result.UpdatedAt,
	})
}

func (h *TaskHandler) UpdateTaskCategory(c *gin.Context) {
	var body request.UpdateTaskCategory
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "your request body is not valid",
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

	userID := c.GetUint("user_id")

	task := domain.Task{
		ID:         uint(taskID),
		CategoryID: body.CategoryID,
		UserID:     userID,
	}

	result, err := h.Service.UpdateTaskCategory(&task)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

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
		UserID:      result.UserID,
		CategoryID:  result.CategoryID,
		UpdatedAt:   result.UpdatedAt,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	urlParam := c.Param("id")
	taskID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	userID := c.GetUint("user_id")

	task := domain.Task{
		ID:     uint(taskID),
		UserID: userID,
	}

	if err := h.Service.DeleteTask(&task); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.Message{
				Message: err.Error(),
			})
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Task has been successfully deleted",
	})
}
