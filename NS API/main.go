package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://gateway.apiportal.ns.nl/reisinformatie-api/api/v2/arrivals?station=rsd", nil)
	request.Header["Ocp-Apim-Subscription-Key"] = []string{"a3df8868b6344301a54e2c2146254794"}
	// q := url.Values{}
	// q.Add("stations", "rsd")
	// request.URL.RawQuery = q.Encode()
	values := request.URL.Query()
	values.Add("stations", "rsd")
	request.URL.RawQuery = values.Encode()
	response, _ := client.Do(request)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
