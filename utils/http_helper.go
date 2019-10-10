package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	httpHelper := NewHttpHelper()
	httpHelper.Get("sdf")
}

type HttpHelper struct {
	client *http.Client
}

func NewHttpHelper() *HttpHelper {
	return &HttpHelper{&http.Client{}}
}

func (h *HttpHelper) Get(url string) interface{} {

	// resp, err := http.Get("https://api.gopax.dev.streami.io/assets")

	req, _ := http.NewRequest("GET", "https://api.gopax.dev.streami.io/assets", nil)
	req.Header.Add("KEY", "!@#$")
	resp, err := h.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return body
}

func (h *HttpHelper) Post(url string, body interface{}) {
	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))

	res, err := h.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	fmt.Println(resBody)

}
