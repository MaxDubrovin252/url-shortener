package handler

import (
	"url-shortener/pkg/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Create(c *gin.Context) {

	var req models.Request

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "cannot decode your request"})

		logrus.Error("cannot decode request")
	}

	id, err := h.services.Create(req.URL, req.Alias)

	if err != nil {
		c.JSON(500, gin.H{"error": "internal"})

		logrus.Error("internal error", err)
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h *Handler) Get(c *gin.Context) {
	alias := c.Param("alias")

	if alias == "" {
		c.JSON(400, gin.H{"error": "alias is empty"})
		return
	}

	url, err := h.services.Get(alias)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})

}

func (h *Handler) Delete(c *gin.Context) {
	alias := c.Param("alias")

	if alias == "" {
		c.JSON(400, gin.H{"error": "alias si empty"})
		return
	}

	if err := h.services.Delete(alias); err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		logrus.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"status": "deleted",
	})

}
