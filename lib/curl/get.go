package curl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	//发送请求获取响应
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	//判断响应状态码
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", res.StatusCode))
	}

	//读取响应实体
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
