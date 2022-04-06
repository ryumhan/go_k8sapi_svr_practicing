package ResourceField

import (
	"net/http"

	ServerPropsType "oncue/apiserver/apilib"

	"github.com/julienschmidt/httprouter"
)

type Resource interface {
	Uri() string
	Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
}
