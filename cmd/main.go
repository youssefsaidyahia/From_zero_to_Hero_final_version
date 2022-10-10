package main

import (
	"fmt"
	"fristTry/internal/adpters/api"
	kafka_installation "fristTry/internal/adpters/kafka-installation"
	_ "github.com/gorilla/mux"
	"net/http"
)

func main() {
	go kafka_installation.Consume()
	api.Routers()
	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", api.Logger())
}
