package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Options struct {
	Method  string      `json:"method"`
	Url     string      `json:"url"`
	Headers http.Header `json:"headers"`
	Body    interface{} `json:"body"`
}

func New[R any](options *Options) (R, error) {
	var result R
	client := http.DefaultClient

	var reqBody []byte

	if options.Body != nil {
		reqBody, _ = json.Marshal(options.Body)
	}

	req, err := http.NewRequest(options.Method, options.Url, bytes.NewBuffer(reqBody))

	if err != nil {
		return result, err
	}

	req.Header = options.Headers

	res, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}
