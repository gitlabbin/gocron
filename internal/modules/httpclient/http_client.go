package httpclient

// http-client

import (
	"bytes"
	"fmt"
	"github.com/ouqiang/gocron/internal/lang"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ResponseWrapper struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func Get(url string, timeout int) ResponseWrapper {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return createRequestError(err)
	}

	return request(req, timeout)
}

func PostParams(url string, params string, timeout int) ResponseWrapper {
	buf := bytes.NewBufferString(params)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	return request(req, timeout)
}

func PostJson(url string, body string, timeout int) ResponseWrapper {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/json")

	return request(req, timeout)
}

func request(req *http.Request, timeout int) ResponseWrapper {
	wrapper := ResponseWrapper{StatusCode: 0, Body: "", Header: make(http.Header)}
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	setRequestHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		// handle timeout err with proper info back
		switch err := err.(type) {
		case net.Error:
			if err.Timeout() {
				wrapper.Body = fmt.Sprintf("net.Error with a Timeout-%s", err.Error())
				return wrapper
			}
		case *url.Error:
			fmt.Println("This is a *url.Error")
			if err, ok := err.Err.(net.Error); ok && err.Timeout() {
				wrapper.Body = fmt.Sprintf("*url.Error with a Timeout-%s", err.Error())
				return wrapper
			}
		}

		// handle other errors without analysis
		wrapper.Body = fmt.Sprintf(lang.Tr("http_request_error"), err.Error())
		return wrapper
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wrapper.Body = fmt.Sprintf(lang.Tr("read_http_response_fail"), err.Error())
		return wrapper
	}
	wrapper.StatusCode = resp.StatusCode
	wrapper.Body = string(body)
	wrapper.Header = resp.Header

	return wrapper
}

func setRequestHeader(req *http.Request) {
	req.Header.Set("User-Agent", "golang/gocron")
}

func createRequestError(err error) ResponseWrapper {
	errorMessage := fmt.Sprintf(lang.Tr("create_http_request_fail"), err.Error())
	return ResponseWrapper{0, errorMessage, make(http.Header)}
}
