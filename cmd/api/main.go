package main

import (
	"github.com/andynador/lego/core/api"
	"net/http"
)

func main() {
	r := api.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("asdf"))
	})

	http.ListenAndServe(":3000", r)
}
