package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	// We'll need to define an upgrader
	// this will require a Read and Write buffer size
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		// We'll need to check the origin of our connection
		// this will allow us to make requests from our React
		// development server to here.
		// For now, we'll do no checking and just allow any connection
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}
