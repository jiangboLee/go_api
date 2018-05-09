package utils

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

//const (
//	GetAccess_token = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&"
//	SendTemplate = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
//)

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		Warning(err, "get有问题")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Warning(err, "get_resp读取body有错")
	}
	//fmt.Println(string(body))
	return string(body)
}

func HttpPost(url string, body []byte) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(body))
	if err != nil {
		Danger(err, "post请求失败")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Warning(err, "post_resp读取body有错")
	}
	return string(b)
}