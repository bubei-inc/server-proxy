package impl

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 接受数据包格式(agent to proxy, and transfer to proxy)
type Packet struct {
	Endpoint string
	Data     interface{}
	Type     string
	Time     time.Time
}

type ManagerCenter struct {
	HttpContacts  map[*http.Request]http.ResponseWriter
	Contacts      map[string]*websocket.Conn
	ServerOutChan chan *Packet
	ClientInChan  chan *Packet
	ServerInChan  chan *Packet
	ClientOutChan chan *Packet
}

var Manager = &ManagerCenter{
	HttpContacts:  make(map[*http.Request]http.ResponseWriter),
	Contacts:      make(map[string]*websocket.Conn, 1000),
	ServerOutChan: make(chan *Packet, 10000),
	ClientInChan:  make(chan *Packet, 10000),
	ServerInChan:  make(chan *Packet, 10000),
	ClientOutChan: make(chan *Packet, 10000),
}
