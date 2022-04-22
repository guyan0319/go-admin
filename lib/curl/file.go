package curl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func HttpPostFiles(uri string, params map[string]string, paramName, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 实例化multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 创建multipart 文件字段,注意paramName要跟服务端一致
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
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
