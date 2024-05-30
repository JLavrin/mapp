package util

import (
	"encoding/json"
	"io"
	"net/http"
)

type Req struct {
	Url    string
	Method string
	Body   io.Reader
}

func Request[T any](payload Req) T {
	var result T
	req, err := http.NewRequest("GET", payload.Url, payload.Body)

	if err != nil {
		println("Error while creating request")
		return result
	}

	req.Method = payload.Method
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		println("Error while sending request")
		return result
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		println("Error while decoding response")
		return result
	}

	return result
}
