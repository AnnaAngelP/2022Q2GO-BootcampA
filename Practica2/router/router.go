package router

import (
	"github.com/gorilla/mux"
	controller "main.com/controller"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/writeCSVFile", controller.APIHandler).Methods("GET")
	r.HandleFunc("/readCSVFile", controller.HandlerReadFile).Methods("GET")
	r.HandleFunc("/searchID", controller.HandlerSearchID).Methods("GET")
}
