package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/config"
	"proxy/impl"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func StartServer() {

	fmt.Println("server start ......")
	conf := config.ConfigVal("config")
	http.HandleFunc(conf.ProxyPath, wsHandler)
	http.HandleFunc(conf.ValidatePath, validateHandler)
	http.ListenAndServe(conf.ProxyPort, nil)

}

func validateHandler(resp http.ResponseWriter, req *http.Request) {
	conf := config.ConfigVal("config")
	var msg []byte
	req.Body.Read(msg)
	fmt.Println(string(msg))
	request, err := http.NewRequest("POST", conf.ValidateSite, req.Body)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err  != nil {
		log.Fatal("validate err: ", err)
	}
	defer response.Body.Close()
	resp.WriteHeader(response.StatusCode)
	message, _:= ioutil.ReadAll(response.Body)
	resp.Write(message)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 完成ws协议的握手操作
	// Upgrade:websocket
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	server := impl.InitServerParam(wsConn)
	server.InitServerConnection()

}

