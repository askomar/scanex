package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", i)

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
