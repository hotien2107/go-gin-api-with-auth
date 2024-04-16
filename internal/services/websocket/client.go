package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID       int64 `json:"id"`
	RoomId   int64 `json:"roomId"`
	SenderId int64 `json:"senderId"`
	Conn     *websocket.Conn
	Message  chan *Message
}

type Message struct {
	Content  string `json:"content"`
	Type     string `json:"type"`
	SenderId int64  `json:"senderId"`
	RoomId   int64  `json:"roomId"`
	Status   string `json:"status"`
}
