package gohttp

import "time"

// ClientConfigOption is type used to modify a HTTPClient instance
type ClientConfigOption func(*HTTPClient)

// WithTimeout is a config option that adds/changes the default timeout
// configured in the HTTPClient instance
func WithTimeout(timeout time.Duration) ClientConfigOption {
	return func(c *HTTPClient) {
		c.client.Timeout = timeout
	}
}
