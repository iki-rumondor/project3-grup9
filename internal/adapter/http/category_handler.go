package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/request"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project3-grup9/internal/application"
	"github.com/iki-rumondor/project3-grup9/internal/domain"
)

type CategoryHandler struct {
	Service *application.CategoryService
}

func NewCategoryHandler(service *application.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		Service: service,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	var body request.Category
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

	category := domain.Category{
		Type: body.Type,
	}

	result, err := h.Service.CreateCategory(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	response := response.CreateCategory{
		ID:        result.ID,
		Type:      result.Type,
		CreatedAt: result.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {

	result, err := h.Service.GetCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var categories = []*response.Category{}

	for _, category := range *result {

		var tasks = []*response.TaskCategory{}

		for _, task := range category.Tasks {
			tasks = append(tasks, &response.TaskCategory{
				ID:          task.ID,
				Title:       task.Title,
				Description: task.Description,
				UserID:      task.UserID,
				CategoryID:  task.CategoryID,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   task.UpdatedAt,
			})
		}

		categories = append(categories, &response.Category{
			ID:           category.ID,
			Type:         category.Type,
			UpdatedAt:    category.UpdatedAt,
			CreatedAt:    category.CreatedAt,
			TaskCategory: tasks,
		})
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var body request.Category
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
	categoryID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	category := domain.Category{
		ID:   uint(categoryID),
		Type: body.Type,
	}

	result, err := h.Service.UpdateCategory(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateCategory{
		ID:        result.ID,
		Type:      result.Type,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {

	urlParam := c.Param("id")
	categoryID, err := strconv.Atoi(urlParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
			Message: "please check the url and ensure it follows the correct format",
		})
		return
	}

	category := domain.Category{
		ID: uint(categoryID),
	}

	if err := h.Service.DeleteCategory(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{
		Message: "Category has been successfully deleted",
	})
}
