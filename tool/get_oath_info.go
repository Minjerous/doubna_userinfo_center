package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOAuthInfo(url string, header map[string]string) (info []byte, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err

	}

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return all, err
	err = res.Body.Close()
	if err != nil {
		fmt.Println("请求GET异常", err)
		return
	}
	return nil, err
}
