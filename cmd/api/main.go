package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// rotas são endpoints
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World 2!"))
	})
	fmt.Println("Servidor conectado à Porta :3000")
	http.ListenAndServe(":3000", r)
}
