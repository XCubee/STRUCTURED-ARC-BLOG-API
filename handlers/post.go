package handlers

import (
	"blog-api/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var posts []models.Post

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	// Decode the request body into the 'post' variable
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the post's fields
	if post.Title == "" || post.Content == "" {
		http.Error(w, "Please provide both title and content", http.StatusBadRequest)
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedPost models.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.Id == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.Id == id {

			posts = append(posts[:i], posts[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}
