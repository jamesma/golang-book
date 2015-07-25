package core

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("error in net.Listen()", err)
		return
	}
	for {
		// accept a TCP connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("error in Listener.Accept()", err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("received:", msg)
	}

	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("error in net.Dial()", err)
		return
	}

	// send the message
	msg := "Hello World"
	fmt.Println("sending message:", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("error in Encoder.Encode()", err)
	}

	c.Close()
}

func ServerDemo() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
