package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":4000")
	if err != nil {
		log.Fatal(err.Error())
	}

	go func(conn net.Conn) {
		reader := bufio.NewReader(conn)
		for {
			l, _, err := reader.ReadLine()
			if err != nil {
				log.Println(fmt.Sprintf("error: %s"), err.Error())
			}

			log.Println(string(l))
		}
	}(conn)

	_, err = conn.Write([]byte(`{"jsonrpc": "2.0", "method": "mining.authorize", "params": ["user", "pass"], "id": 3}`))
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = conn.Write([]byte("\n"))

	log.Println("message sent")
	time.Sleep(time.Second)
}
