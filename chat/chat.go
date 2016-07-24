package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket *websocket.Conn
	send chan *message
	room *room
	userData map[string]interface{}
}

func (cli *client) read() {
	for {
	  var msg *message
		if err := cli.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = cli.userData["name"].(string)
			cli.room.forward <- msg
		} else {
			break
		}
	}
	cli.socket.Close()
}

func ( cli *client ) write() {
	for msg := range cli.send {
		if err := cli.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	cli.socket.Close()
}



