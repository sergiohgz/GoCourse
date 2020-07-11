package main

import (
	"io"
	"log"
	"net"
)

type chat struct {
	server net.Listener
	users  []net.Conn
}

type Chat interface {
	EndChat()
	GetServer() net.Listener
	AddUser(newUser net.Conn)
	SendMessage(message string, sender net.Conn)
}

func InitChat(host string, port string) *chat {
	connectionString := host + ":" + port
	listener, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}

	log.Printf("Chat server started on %s", connectionString)

	return &chat{
		server: listener,
		users:  []net.Conn{},
	}
}

func (c *chat) EndChat() {
	c.server.Close()
}

func (c *chat) GetServer() net.Listener {
	return c.server
}

func (c *chat) AddUser(newUser net.Conn) {
	io.WriteString(newUser, "Bienvenido al chat de GDG Marbella!\n")
	c.users = append(c.users, newUser)
}

func (c *chat) SendMessage(message string, sender net.Conn) {
	for _, user := range c.users {
		if user.RemoteAddr().String() != sender.RemoteAddr().String() {
			io.WriteString(user, message)
		}
	}
}
