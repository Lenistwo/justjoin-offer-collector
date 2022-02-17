package rest

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	RequestTimeoutInSeconds = 30
)

var (
	client http.Client
)

func init() {
	client = http.Client{Timeout: time.Duration(RequestTimeoutInSeconds) * time.Second}
}

func SendGetRequest(url string) []byte {
	response, err := client.Get(url)
	checkError(err)

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	return bytes
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
