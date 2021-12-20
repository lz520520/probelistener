package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	connAddr := os.Args[1]
	sendStr := os.Args[2]
	conn, err := net.DialTimeout("tcp", connAddr, 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	conn.Write([]byte(sendStr))

}
