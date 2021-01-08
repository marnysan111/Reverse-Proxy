package main

import "cpu/write"

func main() {
	ip := "192.168.56.1"
	port := "50000"
	write.Write(ip, port)
}
