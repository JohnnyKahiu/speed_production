package api

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// handle login page
	r.HandleFunc("/login", login.GetHandler).Methods("GET")
	r.HandleFunc("/login", login.PostHandler).Methods("POST")

	return r
}
