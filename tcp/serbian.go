package tcp

import (
	"fmt"
	"go-TCP-Serb/util"
	"log"
	"net"
	"time"
)

func StartTCPSerbian(i int) {
	time.Sleep(5 * time.Second)
	conn, err := net.Dial(util.TYPE, util.HOST+":"+util.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if i%2 == 0 {
		fmt.Fprintln(conn, "PEEPEE")
	} else {
		fmt.Fprintln(conn, "POOPOO")
	}

	// bytes, err := io.ReadAll(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bytes))
}
