package main

import (
	"Comnica_SignIN_task/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.RenderForm)
	http.HandleFunc("/upload", handlers.UploadFile)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
