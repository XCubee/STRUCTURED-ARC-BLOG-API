package routes

import (
	"blog-api/handlers"
	"net/http"
)

// RegisterPostRoutes sets up routing for post-related endpoints
func RegisterPostRoutes() {
	// Collection-level routes
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPost(w, r)
		case http.MethodPost:
			handlers.CreatePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handlers.UpdatePost(w, r)
		case http.MethodDelete:
			handlers.DeletePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
