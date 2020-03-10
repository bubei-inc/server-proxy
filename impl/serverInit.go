package impl

import (
	"context"
	"github.com/gorilla/websocket"
)

type ServerInit struct {
	ss *websocket.Conn
}

func InitServerParam(ss *websocket.Conn) (server *ServerInit) {
	return &ServerInit{ss:ss}
}


func (s *ServerInit) InitServerConnection() (serverConn *ServerInit) {
	ctx, cancle := context.WithCancel(context.Background())
	go ServerConnListen(ctx, s.ss, cancle)
	return
}


