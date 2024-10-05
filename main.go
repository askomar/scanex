package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, results chan<- int) {
	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	var openPorts []int
	for i := 1; i < 1024; i++ {
		if port := <-results; port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
	close(ports)
}
