package httpbot

import (
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func NewDocumentFromResponse(res *http.Response) (*goquery.Document, error) {
	var reader io.Reader = res.Body

	contentType := strings.ToLower(res.Header.Get("Content-Type"))
	if strings.HasSuffix(contentType, "charset=euc-jp") {
		reader = transform.NewReader(reader, japanese.EUCJP.NewDecoder())
	} else if strings.HasSuffix(contentType, "charset=shift_jis") {
		reader = transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
	} else if strings.HasSuffix(contentType, "charset=iso-2022-jp") {
		reader = transform.NewReader(reader, japanese.ISO2022JP.NewDecoder())
	}

	return goquery.NewDocumentFromReader(reader)
}