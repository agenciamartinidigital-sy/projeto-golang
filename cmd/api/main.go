package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	// rotas são endpoints
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		id := r.URL.Query().Get("id")
		if param != "" {
			w.Write([]byte(param + " " + id))
		} else {
			w.Write([]byte("teste"))
		}
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

	fmt.Println("Servidor conectado à Porta :3000")
	http.ListenAndServe(":3000", r)
}
