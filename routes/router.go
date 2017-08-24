package routes

import (
	"github.com/gorilla/mux"
)

var BASE_ROUTER = mux.NewRouter().PathPrefix("/jwt-golang").Subrouter()
var SECURED_ROUTER = mux.NewRouter().PathPrefix("/jwt-golang/secured").Subrouter()