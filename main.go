package main

import (
	"blog-api/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.RegisterPostRoutes()
	fmt.Println("🚀 Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
