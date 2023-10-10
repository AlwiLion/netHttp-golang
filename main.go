package main

import (
	"fmt"
	"net/http"

	"github.com/AlwiLion/routes"
)

func main() {
	fmt.Println("Api Using Standard Net/http")
	routes.SetupRoutes()

	port := 3000
	fmt.Printf("Server running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
