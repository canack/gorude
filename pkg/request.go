package gorude

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func mainRequest(target HttpHeader) error {
	domain := target.Domain
	endpoint := target.Endpoint

	url := domain + endpoint

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	ssl := ""

	if target.HTTPS_enabled {
		ssl = "https://"
	} else {
		ssl = "http://"
	}

	req, err := http.NewRequest(target.Type, ssl+url,
		bytes.NewBuffer([]byte(target.PostData)))

	CheckError(err)

	for k, v := range target.Headers {
		req.Header.Set(k, v)
	}

	response, _ := client.Do(req)

	if true {
		bodyBytes, err := io.ReadAll(response.Body)

		CheckError(err)

		bodyString := string(bodyBytes)
		fmt.Println("Status:", response.StatusCode,
			"Content-Length:", len(bodyString),
			"Payload:", target.Payload)
	}
	defer response.Body.Close()

	return nil
}
