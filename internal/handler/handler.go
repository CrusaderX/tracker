package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Processor interface {
	Handle(message, topic string) error
}

type Handler struct {
	processor Processor
}

func NewDefaultHandler(processor Processor) *Handler {
	return &Handler{
		processor: processor,
	}
}

func (h *Handler) Post(c *gin.Context) {
	h.processor.Handle("test", "test")
	c.JSON(http.StatusOK, gin.H{})

}

func (h *Handler) Get(c *gin.Context) {
	h.processor.Handle("test", "test")
	c.JSON(http.StatusOK, gin.H{})
}
