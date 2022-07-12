package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}

/*
$ time ./tcp-scanner-too-fast
./tcp-scanner-too-fast 0.00s user 0.00s system 90% cpu 0.004 total

================================================================
main goroutine doesn’t know to wait for
the connection to take place. Therefore, the code completes
and exits as soon as the for loop finishes its iterations, which
may be faster than the network exchange of packets between
your code and the target ports. You may not get accurate
results for ports whose packets were still in-flight.
*/
