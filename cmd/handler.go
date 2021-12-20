package main

import (
	"fmt"
	"net"
	"probelistener/pkg/log"
	"sync"
	"time"
)

func handler(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	defer conn.Close()
	buf := make([]byte, len(Args.HeaderFlag))
	bufLen, err := conn.Read(buf)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	if string(buf[:bufLen]) == Args.HeaderFlag {
		log.LogInfo(fmt.Sprintf("probe success. local: %s; remote: %s;", conn.LocalAddr().String(), conn.RemoteAddr().String()))
	}
}

func listen(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	log.LogInfo("start listen: " + port)
	var conn net.Conn
	for {
		conn, err = listener.Accept()
		if err != nil {
			log.LogError(err.Error())
			continue
		}
		log.LogInfo("new Connection from: " + conn.RemoteAddr().String())
		go handler(conn)
	}

}
