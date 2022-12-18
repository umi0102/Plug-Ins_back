package ws

import (
	"Plug-Ins/databases/redisServer"
	"Plug-Ins/routers"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 支持跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
}

// 读取一条消息
func (wsConn *wsConnection) wsRead() map[string]interface{} {

	msgType, data, err := wsConn.wsSocket.ReadMessage()
	if err != nil {
		return nil
	}
	if msgType == websocket.TextMessage {
		str := routers.Byte2Str(data)
		m, err := routers.Str2Map(str)
		if err != nil {
			println("接收到的数据格式错误！")
			return nil
		}
		return m
	}
	return nil
}

func (wsConn *wsConnection) wsWrite(msg []byte) error {
	wsConn.wsSocket.WriteMessage(websocket.TextMessage, msg)
	return nil
}

// 发送存活心跳
func (wsConn *wsConnection) procLoop() {

	// 启动一个goroutine发送心跳
	red := redisServer.RedisDb.Get()
	sub := redis.PubSubConn{Conn: red}
	err := sub.PSubscribe("server_info")
	if err != nil {
		return
	}
	go func() {
		for {
			time.Sleep(30 * time.Second)
			err := wsConn.wsWrite([]byte{
				0000,
			})
			if err != nil {
				wsConn.wsSocket.Close()
				return
			}
		}
	}()
	go func() {
		for {
			msg := sub.Receive()
			switch v := msg.(type) {
			case redis.Message:
				go wsConn.ChannelMessage(v.Channel, v.Data)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Println(v)
				return
			}
		}
	}()
}

func Handler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket·
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	// 初始化wsConn连接
	wsConn := &wsConnection{
		wsSocket: wsSocket,
	}
	// 处理器
	go wsConn.procLoop()
}
