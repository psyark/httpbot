package httpbot

import (
	"io"
	"net/http"
)

const FormURLEncoded = "application/x-www-form-urlencoded"

func ReadAllAndCloseBody(r *http.Response) {
	if r != nil {
		defer r.Body.Close()
		_, _ = io.Copy(io.Discard, r.Body)
	}
}
