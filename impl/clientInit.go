package impl

import (
	"context"
	"github.com/gorilla/websocket"
)

type ClientInit struct {
	ws *websocket.Conn
}

func InitClientParam(ss *websocket.Conn) (client *ClientInit) {
	return &ClientInit{ws: ss}
}

func (c *ClientInit) ClientInitConn() (clientConn *ClientInit) {
	ctx, cancle := context.WithCancel(context.Background())
	go ClientWsListen(ctx, c.ws, cancle)
	return
}
