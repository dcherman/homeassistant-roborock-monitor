package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	queueUrl := os.Getenv("QUEUE_URL")
	hassUrl := os.Getenv("HASS_URL")

	for {
		res, err := http.Get(queueUrl)

		if err != nil {
			log.Printf("Error making HTTP request: %v", err)
		} else {
			cmd, err := ioutil.ReadAll(res.Body)

			if err != nil {
				log.Printf("Error reading response body: %v", err)
			} else if res.StatusCode != 200 {
				log.Printf("Unexpected response from queue svc: %d, %s", res.StatusCode, cmd)
			} else if string(cmd) != "" {
				log.Printf("Executing cmd %s", cmd)

				response, err := http.Post(fmt.Sprintf("%s/api/webhook/%s", hassUrl, string(cmd)), "text/plain", bytes.NewBufferString(""))

				if err != nil {
					log.Printf("Error getting response: %v", err)
				} else {
					log.Printf("Response status code: %v", response.StatusCode)
				}
			}
		}

		time.Sleep(time.Second)
	}
}
