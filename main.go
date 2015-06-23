package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

const (
	indexPath = "/"
	wsPath    = "/ws"
	port      = 8080
)

var connList = []*websocket.Conn{}

func main() {
	run()
}

func run() {
	fmt.Println("run()")

	http.Handle(indexPath, http.FileServer(http.Dir("client")))

	http.HandleFunc(wsPath,
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(handler)}
			s.ServeHTTP(w, req)
		})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handler(conn *websocket.Conn) {
	hErr := websocket.Message.Send(conn, "hello world!")
	if hErr != nil {
		fmt.Println("hello Error: ", hErr)
		return
	}

	connList = append(connList, conn)
LOOP:
	for {
		var message string
		fmt.Println("wait receve")
		rErr := websocket.Message.Receive(conn, &message)
		if rErr != nil {
			fmt.Println("receive error: ", rErr)
			break LOOP
		}

		fmt.Println("message:", message)
		for _, c := range connList {
			sErr := websocket.Message.Send(c, message)
			if sErr != nil {
				fmt.Println("send error: ", sErr)
			}
		}
	}

	fmt.Println("end: ", conn)
}
