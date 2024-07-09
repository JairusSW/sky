package tunnel

import (
	"io"
	"log"
	"net"
)

func Create(from string, to string, protocol string) {

	listener, err := net.Listen(protocol, to)
	if err != nil {
		log.Fatalf("Unable to bind to address %s: %v", to, err)
	}
	defer listener.Close()

	log.Printf("%s <-> %s", from, to)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handle(from, protocol, conn)
	}
}

func handle(from string, protocol string, src net.Conn) {
	dst, err := net.Dial(protocol, from)
	if err != nil {
		log.Fatalf("Unable to connect to destination %s: %v", from, err)
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Println("Error copying from src to dst:", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Println("Error copying from dst to src:", err)
	}
}
