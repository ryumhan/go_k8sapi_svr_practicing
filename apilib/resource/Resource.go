package ResourceField

import (
	"net/http"

	"oncue/apiserver/apilib"

	"github.com/julienschmidt/httprouter"
)

type Resource interface {
	Uri() string
	Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
	Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response
}

//implementation of Get
func Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{405, "", nil}
}

//implementation of Post
func Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{405, "", nil}
}

//implementation of Put
func Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{405, "", nil}
}

//implementation of Delete
func Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) ServerPropsType.Response {
	return ServerPropsType.Response{405, "", nil}
}
