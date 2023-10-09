package main

import (
	"fmt"
	"net/http"

	"github.com/simonharrisco/jobboard_go/internal/routes"
)

func main() {
	router := routes.NewRouter()
	port := 8080
	// this is like a string concatenation to make it ":8080"
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Starting server on port %v\n", addr)

	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
