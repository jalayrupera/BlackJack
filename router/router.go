package router

import (
	"github.com/gorilla/mux"
	"github.com/jalayrupera/toggltest/controller"
)

func DefinedRouters(r *mux.Router) {
	r.HandleFunc("/deck", controller.CreateDeck).Methods("POST")
	r.HandleFunc("/deck/{id}", controller.GetDeck).Methods("GET")
	r.HandleFunc("/deck/{id}", controller.DrawCard).Methods("PUT")
}
