package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Body []byte
	Err  error
}

func (r *Request) Call(method string, url string, body []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		r.Err = err
		return
	}

	response, err := client.Do(req)
	if err != nil {
		r.Err = err
		return
	}

	defer response.Body.Close()
	r.Body, r.Err = ioutil.ReadAll(response.Body)
}
