package main

import (
	"github.com/guisfits/fullcycle/infra/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "route.new-direction", producer)

	for {
		_ = 1
	}
}
