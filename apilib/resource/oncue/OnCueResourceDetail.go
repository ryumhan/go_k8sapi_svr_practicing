package oncueResource

import (
	"net/http"
	ServerPropsType "oncue/apiserver/apilib"

	"github.com/julienschmidt/httprouter"
)

//OnCue/:category
type OnCueResourceDetail struct {
}

func (OnCueResourceDetail) Uri() string {
	return "/oncue/:category"
}

func (OnCueResourceDetail) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	if category := ps.ByName("category"); category != "" {
		return ServerPropsType.Response{200, "", "ServerPropsType.SupportData."}
	}

	return ServerPropsType.Response{400, "", nil}
}

func (OnCueResourceDetail) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	if category := ps.ByName("category"); category != "" {
		return ServerPropsType.Response{200, "", "ServerPropsType.SupportData."}
	}

	return ServerPropsType.Response{400, "", nil}
}

func (OnCueResourceDetail) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	if category := ps.ByName("category"); category != "" {
		return ServerPropsType.Response{200, "", "ServerPropsType.SupportData."}
	}

	return ServerPropsType.Response{400, "", nil}
}

func (OnCueResourceDetail) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{200, "This is K8S-OnCue Rest API Server, This Server Support", "data"}
}
