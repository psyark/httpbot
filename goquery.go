package httpbot

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func NewDocumentFromResponse(res *http.Response) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(transformReader(res.Body, res.Header.Get("Content-Type")))
}

func NewDocumentFromBytes(content []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	contentType := doc.Find(`meta`).FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.ToLower(s.AttrOr("http-equiv", "")) == "content-type"
	}).AttrOr("content", "")

	var reader io.Reader = bytes.NewReader(content)
	if tr := transformReader(reader, contentType); tr != reader {
		return goquery.NewDocumentFromReader(tr)
	} else {
		return doc, nil
	}
}

func transformReader(reader io.Reader, contentType string) io.Reader {
	contentType = strings.ToLower(contentType)
	if strings.HasSuffix(contentType, "charset=euc-jp") {
		return transform.NewReader(reader, japanese.EUCJP.NewDecoder())
	} else if strings.HasSuffix(contentType, "charset=shift_jis") {
		return transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
	} else if strings.HasSuffix(contentType, "charset=iso-2022-jp") {
		return transform.NewReader(reader, japanese.ISO2022JP.NewDecoder())
	}
	return reader
}

func GetFormValues(form *goquery.Selection) url.Values {
	vars := url.Values{}
	form.Find("input").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("type", "") != "checkbox" || s.AttrOr("checked", "") == "checked" {
			if name, ok := s.Attr("name"); ok {
				vars.Add(name, s.AttrOr("value", ""))
			}
		}
	})
	return vars
}
