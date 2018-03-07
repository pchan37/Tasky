package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"tasky/src/lib/templateManager"
	"tasky/src/views"
)

type config struct{
	LayoutPath string
	IncludePath string
}

func loadConfig(filename string) {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := config{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error unpacking config:", err)
	}
	log.Println("Layout path:", config.LayoutPath)
	log.Println("Include path:", config.IncludePath)
	templateManager.SetTemplateConfig(config.LayoutPath, config.IncludePath)
}

func main() {
	loadConfig("config.json")
	templateManager.LoadTemplates()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	views.RegisterStaticViews()
	views.RegisterPublicViews()

	server.ListenAndServe()
}
