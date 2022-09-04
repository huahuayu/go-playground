package http

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
)

func Get(url string, header map[string]string, params map[string]string) (respBody []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	if params != nil && len(params) > 0 {
		for k, v := range params {
			q := req.URL.Query()
			q.Set(k, v)
			req.URL.RawQuery = q.Encode()
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Post(url string, header map[string]string, body []byte) (respBody []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
