package customHTTP

import (
	"github.com/iki-rumondor/project3-grup9/internal/application"
)

type TaskHandler struct {
	Service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// func (h *TaskHandler) CreateTask(c *gin.Context) {

// 	userID := c.GetUint("user_id")

// 	var body request.Task
// 	if err := c.BindJSON(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if _, err := govalidator.ValidateStruct(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	task := domain.Task{
// 		Title:       body.Title,
// 		Description: body.Description,
// 		Category_Id: body.Category_Id,
// 	}

// 	result, err := h.Service.CreateTask(&task)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	response := response.CreateTask{
// 		ID:          result.ID,
// 		Title:       result.Title,
// 		Status:      result.Status,
// 		Description: result.Description,
// 		User_Id:     result.User_Id,
// 		Category_Id: result.Category_Id,
// 		Created_At:  result.Created_At,
// 	}

// 	c.JSON(http.StatusCreated, response)
// }

// func (h *TaskHandler) GetTask(c *gin.Context) {

// 	userID := c.GetUint("user_id")
// 	defer utils.Recovery(c)

// 	result, err := h.Service.GetTask(userID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	var tasks response.Tasks

// 	for _, task := range *result {
// 		tasks.Tasks = append(tasks.Tasks, &response.Task{
// 			ID:          task.ID,
// 			Title:       task.Title,
// 			Status:      task.Status,
// 			Description: task.Description,
// 			User_Id:     task.User_Id,
// 			Category_Id: task.Category_Id,
// 			Created_At:  task.Created_At,
// 			User: response.Users{
// 				ID:        task.Users.ID,
// 				Email:     task.Users.Email,
// 				Full_Name: task.Users.Full_name,
// 			},
// 		})
// 	}

// 	c.JSON(http.StatusOK, tasks.Tasks)
// }

// func (h *TaskHandler) UpdateTask(c *gin.Context) {
// 	var body request.Task
// 	if err := c.BindJSON(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if _, err := govalidator.ValidateStruct(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	urlParam := c.Param("id")
// 	taskID, err := strconv.Atoi(urlParam)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: "please check the url and ensure it follows the correct format",
// 		})
// 		return
// 	}

// 	task := domain.Task{
// 		ID:          uint(taskID),
// 		Title:       body.Title,
// 		Description: body.Description,
// 	}

// 	result, err := h.Service.UpdateTask(&task)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.UpdateTask{
// 		ID:          result.ID,
// 		Title:       result.Title,
// 		Description: result.Description,
// 		Status:      result.Status,
// 		User_Id:     result.User_Id,
// 		Category_Id: result.Category_Id,
// 		Updated_At:  result.Updated_At,
// 	})
// }

// func (h *TaskHandler) UpdateStatusTask(c *gin.Context) {
// 	var body request.Task
// 	if err := c.BindJSON(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	if _, err := govalidator.ValidateStruct(&body); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	urlParam := c.Param("id")
// 	taskID, err := strconv.Atoi(urlParam)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
// 			Message: "please check the url and ensure it follows the correct format",
// 		})
// 		return
// 	}

// 	task := domain.Task{
// 		ID:     uint(taskID),
// 		Status: body.Status,
// 	}

// 	result, err := h.Service.UpdateStatusTask(&task)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.UpdateStatusTask{
// 		ID:          result.ID,
// 		Title:       result.Title,
// 		Description: result.Description,
// 		Status:      result.Status,
// 		User_Id:     result.User_Id,
// 		Category_Id: result.Category_Id,
// 		Updated_At:  result.Updated_At,
// 	})
// }

// func (h *TaskHandler) UpdateCategoryTask(c *gin.Context) {

// }

// func (h *TaskHandler) DeleteTask(c *gin.Context) {

// }
