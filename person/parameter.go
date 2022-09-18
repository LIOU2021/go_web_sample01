package person

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func QueryStringGetList(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get("id")
	level := v.Get("level")
	fmt.Fprintln(w, "id: ", id, "level: ", level)
}

func PathGetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	level := vars["level"]
	fmt.Fprintln(w, "id: ", id, "level: ", level)
}
