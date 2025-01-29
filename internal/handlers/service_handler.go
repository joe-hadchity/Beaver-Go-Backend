package handlers

import (
	"server/internal/models"
	"server/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	service *services.ServiceService
}

func NewServiceHandler(service *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: service}
}

// CreateService godoc
// @Summary Create a new service
// @Tags Services
// @Accept json
// @Produce json
// @Param service body models.Service true "Service data"
// @Success 201 {object} models.Service
// @Failure 400 {object} ErrorResponse
// @Router /services [post]
func (h *ServiceHandler) CreateService(c *gin.Context) {
	var req models.Service
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	service, err := h.service.CreateService(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, service)
}

// GetService godoc
// @Summary Get service details
// @Tags Services
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} models.Service
// @Failure 404 {object} ErrorResponse
// @Router /services/{id} [get]
func (h *ServiceHandler) GetService(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	service, err := h.service.GetService(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, service)
}

// ListServices godoc
// @Summary List all services
// @Tags Services
// @Produce json
// @Success 200 {array} models.Service
// @Router /services [get]
func (h *ServiceHandler) ListServices(c *gin.Context) {
	services, err := h.service.ListServices(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, services)
}

// ListServicesByCategory godoc
// @Summary List services by category
// @Tags Services
// @Produce json
// @Param categoryId path int true "Category ID"
// @Success 200 {array} models.Service
// @Router /categories/{categoryId}/services [get]
func (h *ServiceHandler) ListServicesByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	services, err := h.service.ListServicesByCategory(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, services)
}

// UpdateService godoc
// @Summary Update service details
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Param service body models.Service true "Service data"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /services/{id} [put]
func (h *ServiceHandler) UpdateService(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var req models.Service
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}
	req.ID = id

	if err := h.service.UpdateService(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteService godoc
// @Summary Delete a service
// @Tags Services
// @Param id path int true "Service ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /services/{id} [delete]
func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := h.service.DeleteService(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}