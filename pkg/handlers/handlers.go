package handlers

import (
	"net/http"
)

// DataHandler base struct handler http requests
type DataHandler struct {
}

// NewDataHandler returns a new DataHandler
func NewDataHandler() *DataHandler {
	return &DataHandler{}
}

func (dh *DataHandler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", dh.Test)
	return mux
}

func (dh *DataHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
