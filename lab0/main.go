package main

import (
	"main/handlers"
	"net/http"
)

func main() {
	// TODO: some code goes here
	// Fill out the HomeHandler function in handlers/handlers.go which handles the user's GET request.
	// Start an http server using http.ListenAndServe that handles requests using HomeHandler.
	http.ListenAndServe(":8080", http.HandlerFunc(handlers.HomeHandler))
}
