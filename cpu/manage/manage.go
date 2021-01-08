package manage

import (
	"fmt"
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
