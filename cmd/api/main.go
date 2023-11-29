package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 8080

type application struct {
	Domain string
}

func main() {
	// add config
	var app application

	//read commandline

	//connect to database
	app.Domain = "example.com"
	log.Println("listen and serve on port:", port)

	//running server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
