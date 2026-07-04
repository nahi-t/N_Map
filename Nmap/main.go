package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanport(target string, port int) {

	p := strconv.Itoa(port)
	adress := net.JoinHostPort(target, p)

	fmt.Println(adress)

	conn, err := net.DialTimeout("tcp", adress, 2*time.Second)
	if err != nil {
		fmt.Print("port error", p)
	} else {
		fmt.Println("port open", conn)
	}
}

func main() {
	target := "scanme.nmap.org"

	for port := 1; port <= 1024; port++ {
		go scanport(target, port)
	}
}
