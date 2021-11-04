package main

import (
	"log"
	"net/http"

	"oncue/apiserver/apilib/resource/oncue"
	"oncue/apiserver/apilib/server"

	"github.com/julienschmidt/httprouter"
)

var test bool = false

func main() {
	router := httprouter.New()

	server.AddResource(router, new(oncueResource.OnCueResource))
	server.AddResource(router, new(oncueResource.OnCueResourceDetail))

	log.Fatal(http.ListenAndServe(":8080", router))
}
