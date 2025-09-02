package main

import (
	"benedixx-personalized-agent/src/api"
	"benedixx-personalized-agent/src/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("error while loading config:", err)
	}

	router := api.SetupRouter()
	router.Run(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
}
