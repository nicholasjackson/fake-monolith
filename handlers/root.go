package handlers

import "net/http"

// Root handler for the application
type Root struct{}

// Get implements the http Handler interface
func (ro *Root) Get(rw http.ResponseWriter, r *http.Request) {

}
