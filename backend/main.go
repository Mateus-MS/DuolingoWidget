package main

import (
	"fmt"
	"net/http"

	"github.com/Mateus-MS/DuolingoWidget/routes"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", routes.StreakRoute)

	startServer(router)
}

func startServer(router *http.ServeMux) {
	fmt.Println("Starting HTTP server on port 80")
	if err := http.ListenAndServe(":80", router); err != nil {
		fmt.Println("HTTP server error:", err)
	}
}
