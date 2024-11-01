package user

import (
	"github.com/gorilla/mux"
)

func Rutas(router *mux.Router) {
	router.HandleFunc("/saludo", HelloWorld).Methods("GET")
}
