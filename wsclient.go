package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

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
	go WriteMessage(websocketConn) // Start a goroutine to send messages
	go ReadMessage(websocketConn)  // Start a goroutine to read messages
	time.Sleep(2 * time.Minute)    // Wait for 30 seconds to allow message exchange
	os.Exit(0)                     // Exit the program gracefully
}

func WriteMessage(websocketConn *websocket.Conn) {
	data := ""
	for data != "exit" { // Send a message to the WebSocket server
		fmt.Println("Enter message")
		fmt.Scanln(&data)
		errWrite := websocketConn.WriteMessage(websocket.TextMessage, []byte(data))
		if errWrite != nil {
			fmt.Println("Error sending message:", errWrite)
			return
		}
	}
}

func ReadMessage(websocketConn *websocket.Conn) {
	data := ""
	for data != "exit" { // Read message from WebSocket server
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
}
