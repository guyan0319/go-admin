package curl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpDelete(url string) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", res.StatusCode))
	}
	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return rb, nil
}
