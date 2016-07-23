package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn
	send chan []byte
	room *room
}

func (cli *client) read() {
	for {
		if _, msg, err := cli.socket.ReadMessage(); err == nil {
			cli.room.forward <- msg
		} else {
			break
		}
	}
	cli.socket.Close()
}

func ( cli *client ) write() {
	for msg := range cli.send {
		if err := cli.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	cli.socket.Close()
}



