package person

import (
	"fmt"
	"net/http"
)

func Say4(w http.ResponseWriter, r *http.Request) {
	r2 := Rect{width: 2, height: 3}
	c2 := circle{radius: 10}

	fmt.Fprintln(w, "hello4 : ")
	Measure(r2, w)
	Measure(c2, w)

}
