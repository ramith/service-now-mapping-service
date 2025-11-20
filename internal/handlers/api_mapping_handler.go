package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramith/service-now-mapping-service/internal/models"
	"github.com/ramith/service-now-mapping-service/internal/service"
)

type APIMappingHandler struct {
	service service.APIMappingService
}

func NewAPIMappingHandler(service service.APIMappingService) *APIMappingHandler {
	return &APIMappingHandler{service: service}
}

func (h *APIMappingHandler) CreateMapping(c *gin.Context) {
	var req models.CreateAPIMapping
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapping, err := h.service.CreateMapping(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapping)
}

func (h *APIMappingHandler) GetMappingBySysID(c *gin.Context) {
	sysID := c.Param("sys_id")

	mapping, err := h.service.GetMappingBySysID(sysID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapping)
}

func (h *APIMappingHandler) GetMappingByAPIID(c *gin.Context) {
	apiID := c.Param("api_id")

	mapping, err := h.service.GetMappingByAPIID(apiID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapping)
}

func (h *APIMappingHandler) GetMappingByNameAndVersion(c *gin.Context) {
	var req models.SearchByNameVersion
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapping, err := h.service.GetMappingByNameAndVersion(req.APIName, req.APIVersion)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapping)
}

func (h *APIMappingHandler) GetAllMappings(c *gin.Context) {
	mappings, err := h.service.GetAllMappings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappings)
}

func (h *APIMappingHandler) UpdateMapping(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mapping ID"})
		return
	}

	var req models.UpdateAPIMapping
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mapping, err := h.service.UpdateMapping(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapping)
}

func (h *APIMappingHandler) DeleteMapping(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mapping ID"})
		return
	}

	if err := h.service.DeleteMapping(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *APIMappingHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "servicenow-mapping-service",
	})
}
