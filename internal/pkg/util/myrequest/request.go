package myrequest

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"jpnews/internal/pkg/config"
	"net/http"
	"net/url"
)

func RequestGet(url string) (respContent []byte, err error) {
	c := config.GetConfig()
	if c.UseProxy {
		return ProxyGet(c.ProxyUrl, url)
	}

	return Get(url)
}

func ProxyGet(proxyUrl string, requestUrl string) (respContent []byte, err error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}

	transport := &http.Transport{
		Proxy:           proxy,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	content = bytes.TrimPrefix(content, []byte{239, 187, 191}) // remove BOM \uFEFF

	return content, nil
}

func Get(url string) (respContent []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	content = bytes.TrimPrefix(content, []byte{239, 187, 191}) // remove BOM \uFEFF

	return content, nil
}
