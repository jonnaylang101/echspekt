package echspekt

import (
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type When struct {
	method  string
	target  string
	body    io.Reader
	params  map[string]string
	queries url.Values
	headers http.Header
	router  *echo.Echo
	handler echo.HandlerFunc
}

type RequestOption func(*When)

func WhenIMakeARequest(options ...RequestOption) *When {
	w := &When{
		params:  make(map[string]string),
		queries: make(url.Values),
		headers: make(http.Header),
	}

	for _, ro := range options {
		ro(w)
	}

	return w
}

func (r *When) IShouldRecieve(wantStatusCode int, wantErr bool, responseBodyBinder interface{}, options ...EchspektOption) error {
	es := new(Then)
	es.wantStatusCode = wantStatusCode
	es.wantErr = wantErr
	es.bindObject = responseBodyBinder

	for _, opt := range options {
		opt(es)
	}

	return nil
}

// func (r *When) do() (httptest.ResponseRecorder, error) {
// 	req := httptest.NewRequest(r.method, r.target, r.body)
// 	return nil, nil
// }
