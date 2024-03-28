package tcp

import (
	"fmt"
	"go-TCP-Serb/util"
	"io"
	"log"
	"net"
)

func StartTCPSerbian() {
	conn, err := net.Dial(util.TYPE, util.HOST+":"+util.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	bytes, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
