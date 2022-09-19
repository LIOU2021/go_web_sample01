package person

import (
	"fmt"
	"net/http"
)

func Say3(w http.ResponseWriter, r *http.Request) {
	s1 := Struct1{}
	s1.SetName("Mark123")

	fmt.Fprint(w, "hello3 : ", s1.GetName())
}
