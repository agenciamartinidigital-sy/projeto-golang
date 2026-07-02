package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Product struct {
	ID   int
	Name string
}

func main() {
	// rotas são endpoints
	r := chi.NewRouter()
	// Definidos antes das rotas
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		println("endpoint")
		// param := r.URL.Query().Get("name")
		// id := r.URL.Query().Get("id")
		// if param != "" {
		// 	w.Write([]byte(param + " " + id))
		// } else {
		// 	w.Write([]byte("teste"))
		// }
	})

	r.Get("/{productName}/{productId}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		id := chi.URLParam(r, "productId")
		w.Write([]byte(param + " " + id))
	})
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		obj := map[string]string{"message": "sucess"}
		render.JSON(w, r, obj)
	})
	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})
	r.Put("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	fmt.Println("Servidor conectado à Porta :3000")
	http.ListenAndServe(":3000", r)
}

// // Middleware
// func myMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		println("before")
// 		next.ServeHTTP(w, r)
// 		println("after")
// 	})
// }

// func myMiddleware2(next http.Handler) http.Handler {
// 	startTime := time.Now().String()
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		println("request", r.Method, "url ", r.RequestURI, "tempo", startTime)
// 		next.ServeHTTP(w, r)
// 		println("after 2")
// 	})
// }
