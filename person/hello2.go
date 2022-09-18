package person

import (
	"fmt"
	"net/http"
)

func Say2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello2~")
}
