package handlers

import (
	"server/internal/models"
	"server/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// CreateCategory godoc
// @Summary Create a new service category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category data"
// @Success 201 {object} models.Category
// @Failure 400 {object} ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	category, err := h.service.CreateCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory godoc
// @Summary Get category details
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} ErrorResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	category, err := h.service.GetCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, category)
}

// ListCategories godoc
// @Summary List all categories
// @Tags Categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.service.ListCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary Update category details
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category data"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}
	req.ID = id

	if err := h.service.UpdateCategory(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := h.service.DeleteCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}