package main

import (
	"bufio"
	"log"
)

func main() {
	c := InitChat("localhost", "8888")
	defer c.EndChat()
	server := c.GetServer()
	for {
		user, errUser := server.Accept()
		if errUser != nil {
			log.Printf("failed to accept connection: %s", errUser)
			continue
		}

		c.AddUser(user)

		go func() {
			for {
				msg, errMsg := bufio.NewReader(user).ReadString('\n')
				if errMsg != nil {
					log.Println(errMsg)
					continue
				}

				c.SendMessage(msg, user)
			}
		}()
	}
}
