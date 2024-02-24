package httpbot

import (
	"io"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func NewClient() (*http.Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	return &http.Client{Jar: jar}, nil
}

const FormURLEncoded = "application/x-www-form-urlencoded"

func ReadAllAndCloseBody(r *http.Response) {
	if r != nil {
		defer r.Body.Close()
		_, _ = io.Copy(io.Discard, r.Body)
	}
}
