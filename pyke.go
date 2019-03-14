package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"pyke/handlers"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.IndexHandler)

	http.ListenAndServe(":8000", router)
}
