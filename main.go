package main

import (
	"Funcy/handlers"
	_ "Funcy/mypackage" // Ensure this import to trigger the init() function in mypackage
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/invoke-function", handlers.InvokeFunctionHandler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
