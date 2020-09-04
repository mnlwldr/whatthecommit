package whatthecommit

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const url string = "http://whatthecommit.com/index.txt"

// Return a the string from http://whatthecommit.com/index.txt
func Message() string {

	whatTheCommitClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := whatTheCommitClient.Do(req)
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
	return strings.Trim(string(body), "\n")
}
