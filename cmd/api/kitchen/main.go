package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/supwr/saga-pattern/internal/config"
	"github.com/supwr/saga-pattern/pkg/kafka"
	"log"
	"net/http"
)

func main() {
	api := gin.Default()

	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	producer, err := kafka.NewProducer(kafka.Config{
		Brokers: cfg.KafkaBroker,
		Topic:   cfg.KafkaKitchenTopic,
	})
	if err != nil {
		log.Fatalln(err)
	}

	payload := struct {
		Data string `json:"data"`
	}{
		Data: "hello world!",
	}

	p, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	message := kafka.Message{
		Data: p,
		Key:  []byte("8478ef1e-a2bf-42cd-a157-9e627149ec0c"),
	}

	api.GET("/health", func(ctx *gin.Context) {
		err := producer.SendMessage(message)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.String(http.StatusOK, "Kitchen OK")
		return
	})

	api.Run(":8003")
}
