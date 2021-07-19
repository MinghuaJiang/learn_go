package product_v2

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket(ws *websocket.Conn) {
	done := make(chan struct{})
	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				return
			}

			fmt.Printf("received message %s\n", msg.Data)
		}

		close(done)
	}(ws)
loop:
	for {
		select {
		case <-done:
			fmt.Println("connection was closed, lets break out of here")
			break loop
		default:
			products, err := getTopTenProducts()

			if err != nil {
				log.Println(err)
				break
			}

			if err := websocket.JSON.Send(ws, products); err != nil {
				log.Println(err)
				break
			}

			time.Sleep(10 * time.Second)
		}
	}
}
