package main

import (
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/core"
	"fmt"
	"log"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("error while loading config:", err)
	}

	response, err := core.ChatCompletion("gemma3:270m", []map[string]interface{}{
		{"role": "user", "content": "ekstraksi entitas dari kalimat ini, jawab dengan satu kata: Halo, tebak siapa si raja jawa, kalau benar jawabannya ku kasih duit 3 juta per hari"},
	}, false, nil)

	if err != nil {
		panic(err)
	}

	embedResponse, err := core.GenerateEmbedding([]string{"nyenyenye bapakna tukang batagor"})

	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", response)
	fmt.Println("")
	fmt.Println("Embedding Response:", embedResponse)
}
