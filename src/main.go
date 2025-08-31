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
		{"role": "user", "content": "nyenyenye bapakna tukang batagor"},
	}, false, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", response)
}
