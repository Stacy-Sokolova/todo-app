package handler

import (
	"net/http"
	"strconv"
	"todo-app/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Create(c *gin.Context) {
	var input entity.InsertInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	id, err := h.service.Tasks.Create(c.Request.Context(), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	logrus.Println("Create request")
	c.JSON(http.StatusOK, map[string]interface{}{
		"created id": id,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	result, err := h.service.Tasks.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	var input entity.UpdateInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	affected, err := h.service.Tasks.Update(c.Request.Context(), id, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"updated": affected,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	affected, err := h.service.Tasks.Delete(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"updated": affected,
	})
}
