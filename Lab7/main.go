package main

import (
	"Lab7/packages/handlers"
	"fmt"
	"net/http"
)

func main() {
	run()
}


func run() {

	mux := http.NewServeMux()

	mux.HandleFunc(QUADRATIC_EQUATION_ROUTE, handlers.QuadraticEquation)
	mux.HandleFunc(SLICE_ROUTE, handlers.Slice)
	mux.HandleFunc("/", handlers.Index)

	fmt.Println("Server is launching on " + SERVER_URL + SEVER_PORT)
	err := http.ListenAndServe(SEVER_PORT, mux)
	if err != nil {
		fmt.Println("An error occurred while starting the server", "\n", err.Error())
		return
	}
}
