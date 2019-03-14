package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"msg": "hello, pyke!"}`))
}
