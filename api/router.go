package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/login", login.PostHandler).Methods("POST")

	return r
}
