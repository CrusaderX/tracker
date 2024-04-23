package main

import (
	"github.com/CrusaderX/tracker-api/internal/handler"
	"github.com/CrusaderX/tracker-api/internal/processor"
	"github.com/CrusaderX/tracker-api/internal/producer"
	"github.com/CrusaderX/tracker-api/internal/router"

	"github.com/gin-gonic/gin"
)

const (
	topic = "dsp-tracking-event"
)

func main() {
	// move handler & router into api

	gin.SetMode(gin.ReleaseMode)
	port := 8000

	defaultProducer, err := producer.New([]string{"localhost:9092"})

	if err != nil {
		panic(err)
	}

	defaultProcessor := processor.NewDefaultProcessor(defaultProducer)
	defaultHandler := handler.NewDefaultHandler(defaultProcessor)

	defaultRouter := router.NewDefaultRouter(defaultHandler, port)

	err = defaultRouter.Run()

	if err != nil {
		panic(err)
	}
}
