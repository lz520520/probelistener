package main

import (
	"fmt"
	"net"
	"probelistener/pkg/log"
	"strconv"
	"sync"
	"time"
)

func handler(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	defer conn.Close()
	buf := make([]byte, Args.RecvLen)
	bufLen, err := conn.Read(buf)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	if string(buf[:len(Args.HeaderFlag)]) == Args.HeaderFlag {
		log.LogInfo(fmt.Sprintf("probe success. local: %s; remote: %s; data: %s", conn.LocalAddr().String(), conn.RemoteAddr().String(), string(buf[:bufLen])))
	}
}

func udpHandler(conn net.Conn) {

}

func tcpListen(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	log.LogDebug("start listen: " + port)
	var conn net.Conn
	for {
		conn, err = listener.Accept()
		if err != nil {
			log.LogError(err.Error())
			continue
		}
		log.LogDebug(fmt.Sprintf("new Connection %s from: %s", conn.LocalAddr().String(), conn.RemoteAddr().String()))
		go handler(conn)
	}
}

func udpListen(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	//listenAddr := fmt.Sprintf(":%s", port)
	//listener, err := net.Listen("udp", listenAddr)
	//if err != nil {
	//	log.LogError(err.Error())
	//	return
	//}
	var conn *net.UDPConn
	var err error
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	conn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: portInt,
		Zone: "",
	})
	if err != nil {
		log.LogError(err.Error())
		return
	}
	log.LogDebug("start listen: " + port)
	for {

		buf := make([]byte, Args.RecvLen)
		bufLen, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		if string(buf[:len(Args.HeaderFlag)]) == Args.HeaderFlag {
			log.LogInfo(fmt.Sprintf("probe success. data: %s;  remote: %s;", string(buf[:bufLen]), addr.String()))
		}
	}
}
