package app

import (
	"fmt"
	"../router"
	"log"
	"net/http"
)

func StartApp(){
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
