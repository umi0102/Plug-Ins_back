package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func wsActions(data interface{}) []byte {
	go func() {
		for {
			time.Sleep(1 * time.Second)

		}

	}()
	client := &http.Client{}
	requestBody := fmt.Sprintf("ip=%s", data)
	var jsonStr = []byte(requestBody)
	requst, err1 := http.NewRequest("POST",
		"http://localhost:999/test",
		bytes.NewBuffer(jsonStr))

	if err1 != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	response, err2 := client.Do(requst)

	if err2 != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	body, err3 := io.ReadAll(response.Body)
	if err3 != nil {
		panic(map[string]interface{}{
			"code": 1,
			"msg":  "验证码发送失败，请联系管理员",
		})
	}
	var tempMap map[string]interface{}
	err4 := json.Unmarshal(body, &tempMap)
	if err4 != nil {
		panic("Json错误")
	}
	marshal, _ := json.Marshal(tempMap)
	return marshal
}
