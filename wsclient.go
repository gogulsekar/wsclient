package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("WebSocket client is running...")
	url := url.URL{
		Scheme: "wss",
		Host:   "echo.websocket.org"}
	fmt.Println(url.String())
	websocketConn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket server:", err)
		return
	}
	defer websocketConn.Close()

	errWrite := websocketConn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket server!"))
	if errWrite != nil {
		fmt.Println("Error sending message:", errWrite)
		return
	}

	for { // Read message from WebSocket server
		_, message, err := websocketConn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", message)

		//TODO:- learn how to break the loop
		//time.Sleep(30 * time.Second)
		//websocketConn.Close()
	}

	os.Exit(0) // Exit the program gracefully
}
