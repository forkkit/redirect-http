package function

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	http.Redirect(w, r, "https://www.openfaas.com/", http.StatusTemporaryRedirect)
}
