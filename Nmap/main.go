package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

// 👷 Worker function that pulls ports off the channel conveyor belt
func worker(target string, ports chan int, wg *sync.WaitGroup) {
	for port := range ports {
		p := strconv.Itoa(port)
		address := net.JoinHostPort(target, p)

		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			// Port is closed or timed out, move to the next one
			continue
		}

		// Port is open, clean up the connection and print success
		conn.Close()
		fmt.Printf("Port %s is OPEN! 🎉\n", p)
	}
	// Signal that this specific worker is fully finished
	wg.Done()
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <target>")
		os.Exit(1)
	}
	target := os.Args[1]
	var wg sync.WaitGroup

	// 📦 1. Create the shared conveyor belt channel
	ports := make(chan int, 1024)

	// 🚀 2. Start exactly 100 background workers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(target, ports, &wg)
	}

	// 📥 3. Send all 1,024 port numbers into the channel
	for port := 1; port <= 1024; port++ {
		ports <- port
	}

	// 🛑 4. Close the channel so workers know the job is done
	close(ports)

	// ✋ 5. Wait until all 100 workers finish their remaining tasks
	wg.Wait()
	fmt.Println("Scan complete! 👍")
}
