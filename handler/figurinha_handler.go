package handler

import (
	"errors"
	"mu-pond-go/domain"
	"mu-pond-go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FigurinhaHandler struct {
	service service.FigurinhaService
}

func NewFigurinhaHandler(svc service.FigurinhaService) *FigurinhaHandler {
	return &FigurinhaHandler{service: svc}
}

func (h *FigurinhaHandler) Create(c *gin.Context) {
	var req domain.CreateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	figurinha, err := h.service.Create(req)
	if err != nil {
		c.JSON(serviceErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, figurinha)
}

func (h *FigurinhaHandler) List(c *gin.Context) {
	tipo := c.Query("tipo")
	posicao := c.Query("posicao")

	figurinhas, err := h.service.List(tipo, posicao)
	if err != nil {
		c.JSON(serviceErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, figurinhas)
}

func (h *FigurinhaHandler) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	figurinha, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(serviceErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, figurinha)
}

func (h *FigurinhaHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var req domain.UpdateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	figurinha, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(serviceErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, figurinha)
}

func (h *FigurinhaHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(serviceErrorStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (uint, error) {
	val, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func serviceErrorStatus(err error) int {
	switch {
	case errors.Is(err, service.ErrFigurinhaNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrInvalidTipo),
		errors.Is(err, service.ErrInvalidPosicao),
		errors.Is(err, service.ErrCampoObrigatorio):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
