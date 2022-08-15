package main

import (
	"encoding/hex"
	"probelistener/pkg/log"
	"sync"
)

func main() {
	parser()
	//fmt.Println(strings.ReplaceAll("\\x123", `\x`))
	portSlice := make([]string, 0)
	switch Args.ListenRange {
	case "normal":
		Args.ListenRange = "21-23,80-90,135,137,161,389,443,445,873,1099,1433,1521,1900,2082,2083,2222,2375,2376,2601,2604,3128,3306,3311,3312,3389,4440,4848,5001,5432,5560,5900-5902,6082,6379,7001-7010,7778,8009,8080-8090,8649,8888,9000,9200,10000,11211,27017,28017,50000,51111,50030,50060"
	case "all":
		Args.ListenRange = "1-65535"
	default:
	}
	portSlice = SplitPort(Args.ListenRange)

	switch Args.HeaderFlag {
	case "ldap":
		Args.HeaderFlag = "\x30\x0c\x02\x01\x01"
	case "rmi":
		Args.HeaderFlag = "JRMI\x00"
	default:
	}
	log.LogDebug("header flag: 0x" + hex.EncodeToString([]byte(Args.HeaderFlag)))

	var wg sync.WaitGroup
	for _, port := range portSlice {
		wg.Add(1)
		switch Args.Mode {
		case "tcp":
			go tcpListen(port, &wg)
		case "udp":
			go udpListen(port, &wg)

		}
	}
	wg.Wait()
}
