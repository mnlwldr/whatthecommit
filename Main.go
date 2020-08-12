package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
)

func main() {
	url := "http://whatthecommit.com/index.txt"
	whatThecommitClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := whatThecommitClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Printf("$ git commit -m '%s'\n", strings.Trim(string(body), "\n"))
}
