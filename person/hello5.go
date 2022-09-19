package person

import (
	"errors"
	"fmt"
	"net/http"
)

func Say5(w http.ResponseWriter, r *http.Request) {

	error := errors.New("throw error 123")

	fmt.Fprintln(w, "hello5 : ", error)
}
