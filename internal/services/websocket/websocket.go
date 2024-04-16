package websocket

type SocketHandler struct {
	Pool *Pool
}

func NewSocketHandler(p *Pool) *SocketHandler {
	return &SocketHandler{
		Pool: p,
	}
}
