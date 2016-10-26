package usps

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// URLEncode ...
func URLEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}

// GetRequest ...
func (U *Client) GetRequest(requestURL string) []byte {
	currentURL := ""
	if U.Production {
		currentURL += prodbase
	} else {
		currentURL += devbase
	}
	currentURL += requestURL

	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

// GetRequestHTTPS ...
func (U *Client) GetRequestHTTPS(requestURL string) []byte {
	currentURL := ""
	if U.Production {
		currentURL += prodhttpsbase
	} else {
		currentURL += devhttpsbase
	}
	currentURL += requestURL

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(currentURL)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
