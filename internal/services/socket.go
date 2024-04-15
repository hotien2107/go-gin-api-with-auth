package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SocketServer struct {
	conns     map[*websocket.Conn]bool
	broadcast chan []byte
	upgrader  websocket.Upgrader
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		conns:     make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
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
		fmt.Println(string(msg))
		ws.WriteMessage(messType, []byte("Thanks for the msg: "+string(msg)))
	}
}
