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

	conn.Write([]byte(`{"jsonrpc": "2.0", "method": "mining.authorize", "params": ["user", "pass"], "id": 1}`))
	conn.Write([]byte("\n"))

	conn.Write([]byte(`{"jsonrpc": "2.0", "method": "mining.subscribe", "params": ["cgminer/4.10.0"], "id": 2}`))
	conn.Write([]byte("\n"))

	log.Println("message sent")
	time.Sleep(time.Second)
}
