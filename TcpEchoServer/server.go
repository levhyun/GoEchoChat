package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	PORT := os.Args[1]
	socket, error := net.Listen("tcp", ":"+PORT) // 소켓을 연다
	if nil != error {                            // 에러처리
		log.Println(error)
	} else {
		log.Println("SERVER RUNNIG...\nPORT : " + PORT)
	}
	defer socket.Close() // 메인 함수가 종료되면 소켓이 종료

	for {
		client, error := socket.Accept() // 클라이언트 연결
		if nil != error {                // 에러처리
			log.Println(error)
			continue
		}
		defer client.Close()     // 메인 함수가 종료되면 소켓 연결 해제
		go ClientHandler(client) // 클라이언트 제어 함수 호출
	}
}

func ClientHandler(client net.Conn) {
	receiveBuffer := make([]byte, 4096) // 수신 받을 버퍼
	for {
		length, error := client.Read(receiveBuffer) // 데이터 수신
		if nil != error {                           // 에러처리
			if io.EOF == error {
				log.Println(error)
				return
			}
			log.Println(error)
			return
		}
		if 0 < length {
			data := receiveBuffer[:length]
			log.Println("[Client] " + string(data))
			_, error = client.Write(data[:length]) // 수신 받은 데이터 송신
			if error != nil {                      // 에러처리
				log.Println(error)
				return
			}
		}
	}
}
