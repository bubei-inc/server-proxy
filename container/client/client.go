package client

import (
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/gorilla/websocket"
	"proxy/config"
	"proxy/impl"
	"time"
)

var retryTimes = 1

// 连接客户端，向客户端中转数据
func StartClient() {
	log.Info("client start.....")
	fmt.Println(" client start .....")
	conf := config.ConfigVal("config")
	conn, _, err := websocket.DefaultDialer.Dial(conf.TransferPath, nil)
	if err != nil {
		log.Info("connect to conserver fail: ", err)
		log.Info("start to retry connect ....")
		// 五秒后重试
		time.Sleep(5 * time.Second)
		retryTimes++
		if retryTimes < 10 {
			StartClient()
		}
		log.Info("reconnect conserver failed")
	}
	client := impl.InitClientParam(conn)
	client.ClientInitConn()
}
