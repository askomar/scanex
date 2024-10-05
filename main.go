package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", i)

			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("Port '%d' closed\n", i)
				return
			}
			conn.Close()
			fmt.Printf("Port '%d' opened\n", i)
		}(i)
	}
	wg.Wait()
}
