package curl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpPost(url string, params url.Values, contentType string) ([]byte, error) {
	body := strings.NewReader(params.Encode())
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	res, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", res.StatusCode))
	}
	bodyres, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bodyres, nil
}

func HttpPostForm(requestUrl string, postValue url.Values) ([]byte, error) {
	request, err := http.PostForm(requestUrl, postValue)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()
	if request.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", request.StatusCode))
	}

	rb, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return rb, nil
}

//普通post 二进制提交
func HttpPostByte(requestUrl string, postValue []byte) ([]byte, error) {
	body := bytes.NewReader(postValue)
	request, err := http.NewRequest("POST", requestUrl, body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http.NewRequest,[err=%s][url=%s]", err, requestUrl))
	}
	defer request.Body.Close()
	res, err := http.DefaultClient.Do(request)
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", res.StatusCode))
	}
	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return rb, nil
}
