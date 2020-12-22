package gohttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
)

// RequestConfig is type use to modify a Request instance
type RequestConfig func(*Request)

// RequestBody declares a type for all data that is going to be sent as json
// via POST or PUT requests
type RequestBody map[string]interface{}

func (body *RequestBody) reader() io.Reader {
	// convert map into byte array and later into Reader
	data, _ := json.Marshal(body)
	return bytes.NewReader(data)
}

// WithContext adds a context instance that can be used with the request
func WithContext(ctx context.Context) RequestConfig {
	return func(r *Request) {
		r.ctx = ctx
	}
}

// WithHeader is option configuration type on the Header type
// which appends a key, value pair of extra headers to the
// header instance
func WithHeader(key, value string) RequestConfig {
	return func(r *Request) {
		r.headers = append(r.headers, Header{
			key: key, values: []string{value},
		})
	}
}

// WithAcceptJSONHeader can be used to add application/json mime type to Accept header
func WithAcceptJSONHeader() RequestConfig {
	return WithHeader("Accept", MIMEApplicationJSON)
}

// WithContentTypeJSONHeader can be used to add application/json mime type to Content-Type header
func WithContentTypeJSONHeader() RequestConfig {
	return WithHeader("Content-Type", MIMEApplicationJSON)
}

// WithContentTypeXMLHeader can be used to add application/json mime type to Content-Type header
func WithContentTypeXMLHeader() RequestConfig {
	return WithHeader("Content-Type", MIMETextHTML)
}

// WithAuthorizationTokenHeader can be used to add bearer authorization token
func WithAuthorizationTokenHeader(token string) RequestConfig {
	return WithHeader("Authorization", "Bearer "+token)
}
