package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HarshithRajesh/go-postgres-psql/router"
)

func main(){
	r := router.Router()
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}