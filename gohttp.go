package gohttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient used as an abstraction/wrapper over the default http.Client type
type HTTPClient struct {
	client *http.Client
}

// Header definition of key and values for request/response headers
type Header struct {
	key    string
	values []string
}

// Request defines a type used by HTTPClient as a http request instance
type Request struct {
	method  string
	url     string
	headers []Header
	body    io.Reader
	ctx     context.Context
}

// Method returns the http request verb used
func (req *Request) Method() string {
	return req.method
}

func (req *Request) URL() string {
	return req.url
}

func (req *Request) Headers() []Header {
	return req.headers
}

func (req *Request) Body() io.Reader {
	return req.body
}

// Body defines the response body type
type Body struct {
	io.ReadCloser
}

// Status is the response status of a request
type Status struct {
	Code   int    // e.g 200
	Reason string // e.g "OK"
}

// Response defines a type used by HTTPClient as a http response instance
type Response struct {
	*Request // original request object tied to this response
	Status   Status
	Headers  []Header
	Body
}

// NewHTTPClient builds a new HTTPClient instance with given configuration options
func NewHTTPClient(client *http.Client, options ...ClientConfigOption) *HTTPClient {
	htpclient := &HTTPClient{client: client}

	for _, opt := range options {
		opt(htpclient)
	}
	return htpclient
}

// Get builds a Request instance, applies the given Request configurations and
// performs a GET request to the given url
func (c *HTTPClient) Get(ctx context.Context, url string, options ...RequestConfig) (*Response, error) {
	request := &Request{
		method: http.MethodGet,
		url:    url,
		ctx:    ctx,
	}
	applyOptions(request, options...)

	return c.do(request)
}

// Post builds a Request instance, applies the given Request configurations and
// performs a POST request to the given url
func (c *HTTPClient) Post(ctx context.Context, url string, body io.Reader, options ...RequestConfig) (*Response, error) {
	request := &Request{
		method: http.MethodPost,
		url:    url,
		body:   body,
		ctx:    ctx,
	}
	applyOptions(request, options...)

	return c.do(request)
}

// Put builds a Request instance, applies the given Request configurations and
// performs a PUT request to the given url
func (c *HTTPClient) Put(ctx context.Context, url string, body io.Reader, options ...RequestConfig) (*Response, error) {
	request := &Request{
		method: http.MethodPut,
		url:    url,
		body:   body,
		ctx:    ctx,
	}
	applyOptions(request, options...)

	return c.do(request)
}

// Delete builds a Request instance, applies the given Request configurations and
// performs a DELETE request to the given url
func (c *HTTPClient) Delete(ctx context.Context, url string, body io.Reader, options ...RequestConfig) (*Response, error) {
	request := &Request{
		method: http.MethodDelete,
		url:    url,
		body:   body,
		ctx:    ctx,
	}
	applyOptions(request, options...)

	return c.do(request)
}

// helper func that applies request configurations to Request instance
func applyOptions(request *Request, options ...RequestConfig) {
	for _, opt := range options {
		opt(request)
	}
}

func (c *HTTPClient) do(request *Request) (*Response, error) {

	// check if HTTPClient.client is nil
	if c.client == nil {
		return nil, errors.New("cannot perform request, http client is nil")
	}

	// build http.Request object
	req, err := newRequest(request)
	if err != nil {
		// TODO("Provide a custom error message for error")
		return nil, Error{err.Error(), err}
	}

	// perform request using std lib http.client.Do() method
	res, err := c.client.Do(req)
	if err != nil {
		if er, ok := err.(*url.Error); ok {
			return nil, ConnectionRefusedError{err.Error(), er}
		}
		return nil, Error{err.Error(), err}
	}

	response := Response{
		Request: request,
		Status:  Status{Code: res.StatusCode, Reason: res.Status[4:]},
		Headers: headers(res.Header),
		Body:    Body{ReadCloser: res.Body},
	}
	return &response, nil
}

// this func builds a standard http.Request instance from a htpclient custom
// request instance, it adds any headers and returns
func newRequest(request *Request) (*http.Request, error) {
	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header = buildHeaders(request.headers)

	// run with given context
	if request.ctx != nil {
		req.WithContext(request.ctx)
	}

	return req, nil
}

// IsJSON reads the returned headers and looks
// for the occurrence of 'Content-Type': 'application/json'.
// Returns true if the header exists, false otherwise.
func (r *Response) IsJSON() bool {
	headers := buildHeaders(r.Headers)

	var contains = func(search []string, key string) bool {
		str := strings.Join(search, ",")
		return strings.Contains(str, key)
	}

	return contains(headers["Content-Type"], MIMEApplicationJSON)
}

// IsXML returns true if response headers contains content-type
// text/xml
func (r *Response) IsXML() bool {
	headers := buildHeaders(r.Headers)

	var contains = func(search []string, key string) bool {
		str := strings.Join(search, ",")
		return strings.Contains(str, key)
	}

	return contains(headers["Content-Type"], MIMEApplicationXML)
}

// create map of header from slice of Header
func buildHeaders(headers []Header) map[string][]string {
	if len(headers) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for _, h := range headers {
		m[h.key] = h.values
	}
	return m
}

// create slice of Header from a map
func headers(h map[string][]string) []Header {
	headers := make([]Header, 0, len(h))
	for k, v := range h {
		headers = append(headers, Header{key: k, values: v})
	}
	return headers
}

// DefaultHTTPClient
func DefaultHTTPClient() *HTTPClient {
	c := http.DefaultClient
	c.Timeout = time.Second * 30

	return NewHTTPClient(c)
}

// StatusIsSuccess returns bool value if the response code indicates a success
func StatusIsSuccess(code int) bool {
	return code/100 == http.StatusOK/100
}

func StatusIsClientError(code int) bool {
	return code/100 == http.StatusBadRequest/100
}

func StatusIsServerError(code int) bool {
	return code/100 == http.StatusInternalServerError/100
}

func RaiseForStatus(r *Response) error {
	if !StatusIsSuccess(r.Status.Code) {
		return fmt.Errorf("api request fail, returned status %v", r.Status.Code)
	}
	return nil
}
