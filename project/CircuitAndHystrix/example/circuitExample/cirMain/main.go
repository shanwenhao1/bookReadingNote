package main

import (
	"bookReadingNote/project/CircuitAndHystrix/example/circuitExample/circuitManager"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Get Wraps http.Get in CircuitBreaker
func Get(url string) ([]byte, error) {
	body, err := circuitManager.CirManager.Manager.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {
	body, err := Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
