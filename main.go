package main

import (
	"./controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/reserve", controller.Reserve)
	http.HandleFunc("/cancel", controller.Cancel)
	http.HandleFunc("/query", controller.Query)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
