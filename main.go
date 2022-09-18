package main

import (
	"example/web/person"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	port := ":8000"

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/list/{id}/{level}", person.PathGetList).Methods("GET")
	r.HandleFunc("/list", person.QueryStringGetList).Methods("GET")

	r.HandleFunc("/bye", byeHandler).Methods("GET")
	r.HandleFunc("/t1", person.Say).Methods("POST")
	r.HandleFunc("/t2", person.Say2).Methods("POST")

	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, middleware(r))

}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//You do something here using request
		fmt.Println(r)

		h.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bye world")
}
