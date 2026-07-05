package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)


func worker(target string, ports chan int, wg *sync.WaitGroup) {
	for port := range ports {
		p := strconv.Itoa(port)
		address := net.JoinHostPort(target, p)

		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			
			continue
		}
		conn.Close()
		fmt.Printf("Port %s is OPEN! \n", p)
	}
	
	wg.Done()
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("go run main.go <target> <max-port>")
		os.Exit(1)
	}
	target := os.Args[1]
	po := os.Args[2]
	portInput, err := strconv.Atoi(po)
	if err != nil {
		fmt.Println("Invalid port number:", po)
		os.Exit(1)
	}
	var wg sync.WaitGroup

	
	ports := make(chan int, portInput)

	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(target, ports, &wg)
	}

	for port := 1; port <= portInput; port++ {
		ports <- port
	}
     close(ports)
	wg.Wait()
	fmt.Println("Scan complete! ")
}
