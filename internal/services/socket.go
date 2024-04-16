package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SocketServer struct {
	conns    map[*websocket.Conn]bool
	upgrader websocket.Upgrader
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		conns: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *SocketServer) HandlerWS(ctx *gin.Context) {
	// create a new connection
	conn, err := s.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error: ", err)
		return
	}
	defer conn.Close()
	s.conns[conn] = true
	s.readLoop(conn)
}

func (s *SocketServer) readLoop(ws *websocket.Conn) {
	fmt.Println("Connect to: ", ws.RemoteAddr().String())
	for {
		messType, msg, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read error: ", err)
			continue
		}
		fmt.Println(ws.RemoteAddr().String() + " send message: " + string(msg))
		s.broadcast(messType, msg)
	}
}

func (s *SocketServer) broadcast(msgType int, b []byte) {
	for ws := range s.conns {
		go func() {
			if err := ws.WriteMessage(msgType, b); err != nil {
				fmt.Println("Write error: ", err)
			}
		}()
	}
}
