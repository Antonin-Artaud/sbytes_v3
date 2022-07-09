package commons

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpHelper struct {
	client *http.Client
}

func NewHttpHelper() *HttpHelper {
	return &HttpHelper{
		client: &http.Client{},
	}
}

type IHttpHelper interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) ([]byte, error)
	Put(url string, body []byte) ([]byte, error)
}

func (h *HttpHelper) Get(url string) ([]byte, error) {
	get, err := h.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(get.Body)

	body, err := ioutil.ReadAll(get.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h *HttpHelper) Post(url string, body []byte) ([]byte, error) {
	post, err := h.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(post.Body)

	body, err = ioutil.ReadAll(post.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h *HttpHelper) Put(url string, body []byte) ([]byte, error) {
	put, err := h.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(put.Body)

	body, err = ioutil.ReadAll(put.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}



