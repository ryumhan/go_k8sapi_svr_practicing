package oncueResource

import (
	"net/http"
	ServerPropsType "oncue/apiserver/apilib"

	"github.com/julienschmidt/httprouter"
)

//OnCue
type OnCueResource struct {
}

func (OnCueResource) Uri() string {
	return "/oncue"
}

func (OnCueResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (OnCueResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (OnCueResource) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}

func (OnCueResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}
