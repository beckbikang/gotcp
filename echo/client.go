package main

import (
	"log"
	"net"
	"sync"
	"time"
)

var awg sync.WaitGroup

func main() {

	//建立100个长连接
	for i := 0; i < 100; i++ {
		awg.Add(1)
		go func() {
			doDail()
			awg.Done()
			log.Println("awg end")
		}()
	}
	awg.Wait()
}

func doDail() {
	con, err := net.Dial("tcp", "127.0.0.1:9888")
	if err != nil {
		log.Fatalln("dail 127.0.0.1:9888 error")
	}
	defer con.Close()

	for {
		log.Println("##client send data##")

		con.Write([]byte("start send data"))
		buf := make([]byte, 1024)
		n, err := con.Read(buf)
		if err != nil {
			log.Fatalf("read error %s\n", con.RemoteAddr().String())
			break
		}

		log.Printf("client：read data %s", buf[:n])
		con.SetReadDeadline(time.Now().Add(5 * time.Second))
		log.Println("##client send data##")

		time.Sleep(time.Millisecond * 300)
	}

}
