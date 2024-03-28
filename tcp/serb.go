package tcp

import (
	"bufio"
	"fmt"
	"go-TCP-Serb/util"
	"log"
	"net"
	"time"
)

func StartTCPSerb() {
	listener, err := net.Listen(util.TYPE, util.HOST+":"+util.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Connection timed out.")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "PEEPEE" {
			fmt.Println(line + "POOPOO")
		} else {
			fmt.Println(line + "PEEPEE")
		}
	}

	fmt.Printf("TH'END (ended at %v)\n", time.Now().Format("03:04:05"))
}
