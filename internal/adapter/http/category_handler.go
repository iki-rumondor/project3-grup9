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

type CategoryHandlers struct {
	Service *application.CategoryService
}

func NewCategoryHandler(service *application.CategoryService) *CategoryHandlers {
	return &CategoryHandlers{
		Service: service,
	}
}

func (h *CategoryHandlers) CreateCategory(c *gin.Context) {
	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

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
		ID:         result.ID,
		Type:       result.Type,
		Created_At: result.Created_At,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CategoryHandlers) GetCategorys(c *gin.Context) {
	userID := c.GetUint("user_id")
	defer utils.Recovery(c)

	result, err := h.Service.GetCategorys(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
		return
	}

	var categorys []*response.Category

	for _, category := range *result {
		categorys = append(categorys, &response.Category{
			ID:         category.ID,
			Type:       category.Type,
			Updated_At: category.Updated_At,
			Created_At: category.Created_At,
			Tasks: response.Task{
				ID:          category.Tasks.ID,
				Title:       category.Tasks.Title,
				Description: category.Tasks.Description,
				User_Id:     category.Tasks.User_Id,
				Created_At:  category.Tasks.Created_At,
				Updated_At:  category.Tasks.Updated_At,
			},
		})
	}

	c.JSON(http.StatusOK, categorys)
}

func (h *CategoryHandlers) UpdateCategory(c *gin.Context) {
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
		ID:         result.ID,
		Type:       result.Type,
		Updated_At: result.Updated_At,
	})
}

func (h *CategoryHandlers) DeleteCategory(c *gin.Context) {

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
		Message: "Your Category has been successfully deleted",
	})
}
