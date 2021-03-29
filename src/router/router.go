package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func Routes() {
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{true, "Hello Go!"}
		res, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}).Methods("GET")

	Router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %s!\n", mux.Vars(r)["name"])
	}).Methods("POST")

	Router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := Response{false, "Method not allowed!"}
		res, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(res)
	})

	Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := Response{false, "The page you requested could not be found!"}
		res, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	})
}
