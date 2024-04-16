package websocket

type Room struct {
	ID      int64              `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Pool struct {
	Rooms map[string]*Room `json:"rooms"`
}

func NewPool() *Pool {
	return &Pool{
		Rooms: make(map[string]*Room),
	}
}
