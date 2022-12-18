package ws

import (
	"Plug-Ins/databases/mysql"
	"Plug-Ins/databases/redisServer"
	"Plug-Ins/routers/tools"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// ChannelMessage redis订阅通讯
func (wsConn *wsConnection) ChannelMessage(channel string, data []byte) {
	if channel == "server_info" {
		wsConn.wsWrite(data)
	}
}

type DataList struct {
	code      string
	ArrayList []interface{}
}

// GetserveInfos 发布serve_info订阅
func GetserveInfos() {
	get := redisServer.RedisDb.Get()
	go func() {
		for {
			time.Sleep(6 * time.Second)
			ArrayList := make([]interface{}, 0)
			res := mysql.SelectAllMysql("select * from  kits_server")
			var wg sync.WaitGroup
			// 增加需要等待的任务数
			wg.Add(len(res))
			fmt.Println(len(res))
			for i := 0; i < len(res); i++ {
				go func(i int) {
					re := tools.GetServerInfo(
						res[i]["server_ip"].(string),
						res[i]["server_account"].(string),
						res[i]["server_pwd"].(string),
						res[i]["server_port"].(string))
					ArrayList = append(ArrayList, re)
					defer func() {
						fmt.Println("协程", i, "已完成")
						wg.Done()
					}()
				}(i)

			}
			fmt.Println("等待")
			wg.Wait()
			fmt.Println("完成")

			fmt.Println(ArrayList)
			red := redisServer.RedisDb.Get()
			//fmt.Println(ArrayList)
			ress, _ := M2S(ArrayList)

			_, err1 := red.Do("PUBLISH", "server_info", ress)
			if err1 != nil {
				fmt.Println("server_info管道建立失败！")
			}
			redisServer.SetRedis("server_info_task", "1", 30, get)
		}
	}()
}
func M2S(mapData []interface{}) (result string, err error) {
	resultByte, errError := json.Marshal(mapData)
	result = string(resultByte)
	err = errError
	return result, err
}
