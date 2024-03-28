package main

import "go-TCP-Serb/tcp"

func main() {
	for i := 0; i < 5; i++ {
		go tcp.StartTCPSerbian(i)
	}

	tcp.StartTCPSerb()
}
