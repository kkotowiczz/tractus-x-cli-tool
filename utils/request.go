package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendPostRequest(body []byte, url string, apiKey string) string {
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-API-KEY", apiKey)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("no body")
	}

	defer res.Body.Close()

	return string(bodyBytes)
}
