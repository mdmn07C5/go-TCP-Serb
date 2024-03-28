package serbs

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func StartTCPSerb() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nhenlo?\n")
		fmt.Fprintln(conn, "PEEPEE")
		fmt.Fprintf(conn, "%v", "POOPOO")
		conn.Close()
	}
}
