package zhttp

import (
	"AWork-3/zfunc"
	"AWork-3/zvar"
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func Zhttp1(method, url, payload string) (string, error) {
	//第一阶段，定义请求
	request, err := http.NewRequest(method, url, bytes.NewBufferString(payload))
	if err != nil {
		log.Println("requests", err)
		return "", err
	}
	//1.2设置请求头
	for key, value := range zvar.Headers {
		request.Header.Set(key, value)
	}
	//1.3 发起请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("client", err)
		time.Sleep(6 * time.Second)
		return "", err
	} else {
		log.Println("Response1 Status:", response.Status)
	}

	//1.4读取工单系统工单摘要内容
	body, _ := io.ReadAll(response.Body)
	response.Body.Close()
	//1.5 返回值，工单系统摘要body,这个内容交给
	return string(body), err
}

func InfoMap(body string) []map[string]string {
	newbody := strings.Split(body, "{\"assignee\":")
	var form []map[string]string
	var entry map[string]string
	for i := 1; i < len(newbody); i++ {
		//fmt.Println(i)
		//fmt.Println(newbody[i])

		pattern := `"taskId":"([a-zA-Z0-9]+)","processInstanceId":"([a-zA-Z0-9]+)".*?"processTitle":"(.*?)","processKey":"(.*?)"`
		// 使用正则表达式查找匹配的内容
		matches := zfunc.MatchString(pattern, newbody[i])
		for _, match := range matches {
			entry = map[string]string{
				"taskId":            match[1],
				"processInstanceId": match[2],
				"processTitle":      match[3],
				"processKey":        match[4],
			}
		}
		//将循环的每一次map数据存入到切片中
		form = append(form, entry)
	}
	return form
}
