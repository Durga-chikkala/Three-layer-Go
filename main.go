package main

import (
	"net/http"
	"threeLayer/handlers"
	"threeLayer/services"
	"threeLayer/stores"
)

func main() {
	store := stores.New()
	service := services.New(store)
	handler := handlers.New(service)

	http.HandleFunc("/student", handler.PostMarks)

	http.ListenAndServe(":8000", nil)
}
