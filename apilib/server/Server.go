package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	ServerPropsType "oncue/apiserver/apilib"
	ResourceField "oncue/apiserver/apilib/resource"

	"github.com/julienschmidt/httprouter"
)

func Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func AddResource(router *httprouter.Router, resource ResourceField.Resource) {
	fmt.Println("\"" + resource.Uri() + "\" api is registerd")

	router.GET(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Get(rw, r, ps)
		HttpResponse(rw, r, res)
	})
	router.POST(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Post(rw, r, ps)
		HttpResponse(rw, r, res)
	})
	router.PUT(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Put(rw, r, ps)
		HttpResponse(rw, r, res)
	})
	router.DELETE(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Delete(rw, r, ps)
		HttpResponse(rw, r, res)
	})
}

func GetResource(category string) ServerPropsType.Response {
	return ServerPropsType.Response{}
}

func HttpResponse(rw http.ResponseWriter, req *http.Request, res ServerPropsType.Response) {
	content, err := json.Marshal(res)

	if err != nil {
		Abort(rw, 500)
	}

	rw.WriteHeader(res.Code)
	rw.Write(content)
}
