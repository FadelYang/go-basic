package main

import (
	"go-basic-alex-mux/cmd/tutorial-12/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/items", handler.GetItems)

	http.ListenAndServe(":5000", nil)
}
