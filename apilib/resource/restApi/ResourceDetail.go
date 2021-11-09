package restApi

import (
	"net/http"
	ServerPropsType "oncue/apiserver/apilib"

	"oncue/apiserver/apilib/k8sApi"
	"github.com/julienschmidt/httprouter"
)

//OnCue/:category
type ResourceDetail struct {
}

func (ResourceDetail) Uri() string {
	return "/oncue/:category"
}

func (ResourceDetail) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	if category := ps.ByName("category"); category != "" {
		return k8sApi.GetApi(category)
	}

	return ServerPropsType.Response{400, "", nil}
}

func (ResourceDetail) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	if category := ps.ByName("category"); category != "" {
		return ServerPropsType.Response{200, "", "ServerPropsType.SupportData."}
	}

	return ServerPropsType.Response{400, "", nil}
}

func (ResourceDetail) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{400, "Invalid Request", nil}
}

func (ResourceDetail) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{400, "Invalid Request", nil}
}
