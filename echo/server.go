package main

import (
	"io"
	"log"
	"net"
	"time"
)

const (
	DEFAULT_SLEEP    = 1 * time.Second
	DEFAULT_BUF_SIZE = 4096
)

func main() {

	tcpCon, err := net.Listen("tcp", "127.0.0.1:9888")

	if err != nil {
		log.Fatal("listen 9888 faild")
	}
	for {
		acceptCon, err := tcpCon.Accept()
		if err != nil {
			time.Sleep(DEFAULT_SLEEP)
		}

		go Echo(acceptCon)
	}
}

func Echo(con net.Conn) {
	defer con.Close()
	buf := make([]byte, DEFAULT_BUF_SIZE)

	for {
		log.Println("#####服务端start running")
		n, err := con.Read(buf)
		log.Printf("服务发送数据:%s\n", buf[:n])
		con.Write(buf[:n])

		log.Println("#####服务端end running")
		if err == io.EOF {
			break
		}
	}

}
