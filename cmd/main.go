package main

import (
	"benedixx-personalized-agent/src/api"
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/database"
	"os"
	"os/signal"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic("failed to load config: " + err.Error())
	}
	database.InitDB()
	defer database.CloseDB()

	router := api.SetupRouter()
	router.Run(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	config.Log.Info("Shutting down server...")
}
