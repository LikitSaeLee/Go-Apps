package main

import (
	"net/http"
	"log"
	"trace"
	"github.com/gorilla/websocket"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients
	forward chan []byte

	// join is a channel for client to join the room
	join chan *client

	// leave is a channel for client to leave the room
	leave chan *client

	// client holds all current clients in this room
	clients map[*client]bool
	// tracer will receive trace information of activity
	// in the room
	tracer trace.Tracer
}

func newRoom() *room {
	return &room{
		forward:  make(chan []byte),
		join:     make(chan *client),
		leave:    make(chan *client),
		clients:  make(map[*client]bool),
		tracer:   trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <- r.join:
		  r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <- r.leave:
		  r.clients[client] = false
			r.tracer.Trace("Client left")
		case msg := <- r.forward:
		  // forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
					r.tracer.Trace(" -- sent to client")
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- falied to send, cleaned up client")
				}
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}