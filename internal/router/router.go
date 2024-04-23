package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}

type Router struct {
	r    *gin.Engine
	h    Handler
	port int
}

func NewDefaultRouter(h Handler, port int) *Router {
	r := &Router{
		r:    gin.Default(),
		h:    h,
		port: port,
	}
	r.ConfigureRouter()

	return r
}

func (r *Router) ConfigureRouter() {
	r.r.POST("/events/custom", r.h.Post)
	r.r.GET("/events", r.h.Get)
}

func (r *Router) Run() error {

	if err := r.r.Run(fmt.Sprintf(":%d", r.port)); err != nil {
		return err
	}

	return nil
}
