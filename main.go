package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

const (
	path = "/"
	port = 8080
)

func main() {
	run()
}

func run() {
	fmt.Println("websocket run")
	http.HandleFunc(path,
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
	for {
		var message string
		fmt.Println("wait receve")
		rErr := websocket.Message.Receive(conn, &message)
		if rErr != nil {
			fmt.Println(rErr)
			break
		}

		fmt.Println("message:", message)
		sErr := websocket.Message.Send(conn, message)
		if sErr != nil {
			fmt.Println(sErr)
			break
		}
	}
}
