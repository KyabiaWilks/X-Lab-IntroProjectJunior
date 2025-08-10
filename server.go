package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setCorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func startServer() {
	r := mux.NewRouter()
	r.Use(setCorsHeaders)

	r.HandleFunc("/comment/get", getComments).Methods("GET", "OPTIONS")
	r.HandleFunc("/comment/add", addComment).Methods("POST", "OPTIONS")
	r.HandleFunc("/comment/delete", deleteComment).Methods("POST", "OPTIONS")
	r.HandleFunc("/order/switch", switchOrder).Methods("GET", "OPTIONS")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./static/comment.html")
		} else {
			http.NotFound(w, r)
		}
	})
	invertOrder = true
	http.ListenAndServe(":8080", r)
	//log.Println("Server listening on port 8080...")
}
