package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/config"
	"main.go/controller"
	"main.go/storage/postgres"
)

func main() {
	cfg := config.Load()
	s, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer s.Close()

	con := controller.New(s)
	http.HandleFunc("/users", con.Users)
	fmt.Println("Listening at port 8080...")
	http.ListenAndServe(":8080", nil)
	// http.HandleFunc("/products", con.Users)
	// http.HandleFunc("/basketproducts", con.Users)
	// http.HandleFunc("/category", con.Users)
	// http.HandleFunc("/basket", con.Users)
	

}
