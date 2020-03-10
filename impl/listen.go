package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func ServerConnListen(ctx context.Context, ss *websocket.Conn, stopFun func()) {
	for {
		select {
		case <-ctx.Done():
			if stopFun != nil {
				stopFun()
			}
		case msg := <-Manager.ServerOutChan:
			conn := getConnByKey(msg.Endpoint)
			if conn != nil {
				write(conn, msg)
			}
		default:
			var pac = new(Packet)
			err := ss.ReadJSON(pac)
			fmt.Println(pac)
			if err != nil {
				return
			}
			// send to read chan
			rememberConn(ss, pac)
			Manager.ClientOutChan <- pac

		}
	}
}

func ClientWsListen(ctx context.Context, ss *websocket.Conn, stopFunc func()) {
	for {
		select {
		case <-ctx.Done():
			if stopFunc != nil {
				stopFunc()
			}
			return
		// 写到transfer
		case message := <-Manager.ClientOutChan:
			if ss != nil {
				write(ss, message)
			}
		default:
			var pac = new(Packet)
			err := ss.ReadJSON(pac)
			if err != nil {
				return
			}
			Manager.ServerOutChan <- pac
		}
	}
}

func write(ss *websocket.Conn, message *Packet) {
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(message)
	err := ss.WriteJSON(message)
	if err != nil {
		log.Fatal("send message err, ", err)
	}
}

func rememberConn(ss *websocket.Conn, message *Packet) (conn *websocket.Conn) {
	conn = Manager.Contacts[message.Endpoint]
	if conn != nil {
		return
	}
	Manager.Contacts[message.Endpoint] = ss
	return ss
}

func getConnByKey(key string) (conn *websocket.Conn) {
	return Manager.Contacts[key]
}
