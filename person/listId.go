package person

import (
	"fmt"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	id := v.Get("id")
	level := v.Get("level")
	fmt.Fprintln(w, "id: ", id, "level: ", level)
}
