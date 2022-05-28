package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	router "main.com/router"
)

func main() {
	r := mux.NewRouter()
	router.Routes(r)
	fmt.Println("The server is running...")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8090",
		// Good practice: enforce timeouts for servers you create!
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
