package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kunanon-c/url-shortner/internal/repository"
	"net/http"
)

type Handler struct {
	Repo repository.Repository
}

func (h *Handler) SaveURL(c *gin.Context) {
	var body createShortenURLRequest
	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		return
	}

	shorten, err := h.Repo.Save(body.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "saved",
		"save-to": "localhost:8080/goto/" + shorten,
	})
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
