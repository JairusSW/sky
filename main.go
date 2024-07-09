package main

import "main/tunnel"

func main() {
	tunnel.Create("127.0.0.1:3000", "127.0.0.1:4000", "tcp")
}