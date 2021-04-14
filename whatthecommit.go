package whatthecommit

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const url string = "http://whatthecommit.com/index.txt"

// Message Return a string from http://whatthecommit.com/index.txt
func Message() string {

	whatTheCommitClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkError(err)

	res, getErr := whatTheCommitClient.Do(req)
	checkError(getErr)

	if res.StatusCode != 200 {
		checkError(errors.New(res.Status))
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			checkError(err)
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	checkError(readErr)

	return strings.Trim(string(body), "\n")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
