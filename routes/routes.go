package routes

import (
	controller "bargerbites/controller"

	"github.com/gorilla/mux"
)

func RouterDefination(r *mux.Router) {

	r.HandleFunc("/signup", controller.SignUpUser).Methods("POST")
	r.HandleFunc("/hello", controller.HomeLink).Methods("POST")

}
