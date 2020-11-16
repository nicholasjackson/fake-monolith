package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/nicholasjackson/fake-monolith/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "bind address for server, i.e. :9090")

func main() {

	err := env.Parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	l := hclog.Default()

	hh := &handlers.Root{}
	ph := handlers.NewProduct(l)

	r := mux.NewRouter()
	r.HandleFunc("/", hh.Get).Methods(http.MethodGet)
	r.HandleFunc("/products", ph.Get).Methods(http.MethodGet)
	r.HandleFunc("/products/{id:[0-9]+}", ph.Get).Methods(http.MethodGet)

	http.Handle("/", r)

	http.ListenAndServe(*bindAddress, nil)
}
