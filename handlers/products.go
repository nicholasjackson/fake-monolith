package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Products handler for the application
type Products struct {
	logger hclog.Logger
	Data   []Product `fakesize:"10"`
}

func NewProduct(l hclog.Logger) *Products {
	p := &Products{logger: l}
	gofakeit.Struct(p)

	return p
}

// Get implements the http Handler interface
func (p *Products) Get(rw http.ResponseWriter, r *http.Request) {
	p.logger.Info("Handle products request")

	// get id if present
	vars := mux.Vars(r)
	if id := vars["id"]; id != "" {
		prod := p.getProduct(id)

		if p == nil {
			http.Error(rw, "Product does not exist", http.StatusNotFound)
			return
		}

		e := json.NewEncoder(rw)
		e.Encode(prod)

		return
	}

	// else return everthing
	e := json.NewEncoder(rw)
	e.Encode(p.Data)
}

func (p *Products) getProduct(id string) *Product {
	nId, _ := strconv.Atoi(id)
	for _, prod := range p.Data {
		if prod.ID == nId {
			return &prod
		}
	}

	return nil
}

type Product struct {
	ID               int    `fake:"{number:1,1234}"`
	Name             string `fake:"{hipsterword}"`
	ShortDescription string `fake:"{hipstersentence:10}"`
	LongDescription  string `fake:"{paragraph:1,4,30}"`
	PrimaryColor     string `fake:"{color}"`
}
