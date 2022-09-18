package person

import (
	"fmt"
	"net/http"
)

var Name string = "小明"

func Say(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello1~")
}
