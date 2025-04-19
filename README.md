#  Simple Blog API in Go (Pure net/http)

This project is a simple **Blog API built entirely using Go’s `net/http` package** — no third-party routing frameworks like Gorilla Mux or Gin. It's designed for **beginners who want a deep understanding of how Go’s HTTP server works under the hood**.

---

## Features

-  Pure `net/http` implementation (no external router)
-  RESTful API endpoints: - `GET /posts` - List all blog posts
- `POST /posts` - Create a new blog post
- `PUT /posts/{id}` - Update a blog post by ID
- `DELETE /posts/{id}` - Delete a blog post by ID
-  In-memory storage using slices (no DB for simplicity)
-  Educational and beginner-friendly
-  Project structure that mimics real-world modular architecture

---

## 🛠 Tech Stack

- Language: [Go (Golang)](https://golang.org/)
- HTTP: Native `net/http` package
- Encoding: Standard `encoding/json`

---

## Folder Structure
blog-api/
│
├── main.go                  
│
├── models/
│   └── post.go            
│
├── handlers/
│   └── post.go              
│
├── routes/
│   └── post.go              
│
└── README.md                






## 📦 How to Run

1. **Clone the repo**:
   ```bash
   git clone https://github.com/your-username/blog-api-go.git
   cd blog-api-go
   

