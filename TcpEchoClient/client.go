package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	PORT := os.Args[1]
	client, error := net.Dial("tcp", ":"+PORT) // 소켓 연결
	if nil != error {                          // 에러 처리
		log.Println(error)
	} else {
		log.Println("SERVER CONNECTING...\nPORT : " + PORT)
	}

	go func() {
		receiveBuffer := make([]byte, 4096) // 수신 받을 버퍼

		for {
			length, error := client.Read(receiveBuffer) // 데이터 수신
			if error != nil {                           // 에러 처리
				log.Println(error)
				return
			}

			log.Println("[Server] " + string(receiveBuffer[:length]))
		}
	}()

	for {
		var sendBuffer string // 송신 데이터 변수(
		fmt.Print("> ")
		fmt.Scanln(&sendBuffer)
		client.Write([]byte(sendBuffer)) // 데이터 송신
		time.Sleep(time.Duration(5) * time.Millisecond)
	}
}
