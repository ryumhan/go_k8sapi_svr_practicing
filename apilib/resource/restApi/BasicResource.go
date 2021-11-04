package restApi

import (
	"net/http"
	ServerPropsType "oncue/apiserver/apilib"

	"github.com/julienschmidt/httprouter"
)

//OnCue
type BasicResource struct {
}

func (BasicResource) Uri() string {
	return "/oncue"
}

func (BasicResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (BasicResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (BasicResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (BasicResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}
