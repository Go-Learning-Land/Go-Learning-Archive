package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	dst, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln("Unable to connect to localhiost:3000")
	}
	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, conn); err != nil {
			log.Fatalln("Unable to copy data")
		}
	}()
	if _, err := io.Copy(conn, dst); err != nil {
		log.Fatalln("Unable to copy data")
	}

}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
