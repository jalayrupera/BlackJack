package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jalayrupera/toggltest/router"
)

func main() {
	fmt.Println("Decks Proto")

	r := mux.NewRouter()
	fmt.Println("Server initialized")
	router.DefinedRouters(r)
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening at port 8000...")
}
