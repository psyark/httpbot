package httpbot

import "net/http"

type ResponseChainer func(*http.Response) (*http.Response, error)

func RunChain(chain []ResponseChainer) error {
	var res *http.Response
	var err error
	for _, c := range chain {
		res, err = c(res)
		if err != nil {
			return err
		}
	}
	return nil
}
