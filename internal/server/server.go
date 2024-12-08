package server

import "net/http"

func NewServer(port string, blogHandler BlogHandler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/posts", blogHandler.GetAll)
	mux.HandleFunc("POST /api/v1/posts", blogHandler.Create)
	mux.HandleFunc("GET /api/v1/posts/{id}", blogHandler.GetByID)

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
