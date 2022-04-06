package main

import (
	"log"
	"net/http"

	"oncue/apiserver/apilib/resource/restApi"
	"oncue/apiserver/apilib/server"

	"github.com/julienschmidt/httprouter"
)

var test bool = false

func main() {
	router := httprouter.New()	

	server.AddResource(router, new(restApi.BasicResource))
	server.AddResource(router, new(restApi.ResourceDetail))

	log.Fatal(http.ListenAndServe(":8080", router))
}
