package htpclient

import "time"

//ClientConfigOption is type used to modify a HtpClient instance
type ClientConfigOption func(*HtpClient)

//WithTimeout is a config option that adds/changes the default timeout
//configured in the HtpClient instance
func WithTimeout(timeout time.Duration) ClientConfigOption {
	return func(c *HtpClient) {
		c.client.Timeout = timeout
	}
}
