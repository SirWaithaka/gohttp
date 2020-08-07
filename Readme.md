# HTPCLIENT

This is a simple abstraction library over the Golang std lib `http.Client` type.

## Usage
The library uses a `Request` struct which is a wrapper around `http.Request` type. The `Request` type can be used with
extra configuration options following the `RequestConfig` type.

A sample `GET` request would look like the following.

```go
package main

import (
    "log"
    "net/http"
)

const url = "https://jsonbox.io/demobox_6d9e326c183fde7b"

func main() {

    client := htpclient.NewHtpClient(http.DefaultClient, htpclient.WithTimeout(30*time.Second))    

    // we are going to tell the api that we accept json as valid response
    options := []htpclient.RequestConfig{htpclient.WithAcceptJSONHeader()}
    response, err := client.Get(url, options...)
    if err != nil {
        log.Println(err)
        return
    }
    log.Println(response)
}
```

A sample `POST` request would look like the following.

```go
package main

import (
    "bytes"
    "encoding/json"
    "log"
    "time"
)

const url = "https://jsonbox.io/box_45bedf7b0a8c89ca223a"

func main() {
    client := htpclient.NewHtpClient(http.DefaultClient, htpclient.WithTimeout(30*time.Second))

	data := map[string]interface{}{
		"city":         "Nairobi",
		"country":      "Kenya",
		"country_code": "KE",
		"currency":     "KES",
	}

	// convert data payload into bytes
	var body bytes.Buffer
	_ = json.NewEncoder(&body).Encode(&data)

	// we will pass this options to the request for extra configuration before the request is sent
	options := []htpclient.RequestConfig{htpclient.WithAcceptJSONHeader(), htpclient.WithContentTypeJSONHeader()}
	response, err := client.Post(url+"/cities", bytes.NewReader(body.Bytes()), options...)
	if err != nil {
		log.Println(err)
		return
	}
    log.Println(response)
}
```