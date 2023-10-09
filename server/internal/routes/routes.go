package routes

import (
	"fmt"
	"net/http"
)

func Test() {
	fmt.Println("Hello world")
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
