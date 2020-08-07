package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sirwaithaka/htpclient"
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
	res, err := client.Post(url+"/cities", bytes.NewReader(body.Bytes()), options...)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Status)   // get the status code from the response
	log.Println(res.Headers)  // get the headers returned from the response
	log.Println(res.IsJSON()) // returns true if the response headers reports "Content-Type: application/json"
	log.Println(res.Request)  // the request instance tied to this response
	log.Println(res.Method()) // the http verb used to perform the request

	// perform an example get request to the same endpoint and get our city back
	// the endpoint returns an array of cities
	res, err = client.Get(url+"/cities", options...)
	if err != nil {
		log.Println(err)
		return
	}

	// create a slice type we can store the response
	var cities []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&cities)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(cities) // print the response
}
