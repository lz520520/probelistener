package main

import (
	"encoding/hex"
	"flag"
	"os"
	"probelistener/pkg/log"
	"strconv"
	"strings"
)

var Args *Options

type Options struct {
	ListenRange string
	HeaderFlag  string
}

func SplitPort(ports string) (portSlice []string) {
	// 分割端口列表，输出单一端口切片
	tmpSlice := strings.Split(ports, ",")
	for _, i := range tmpSlice {
		i = strings.TrimSpace(i)
		if i != "" {
			if strings.Contains(i, "-") {
				iSlice := strings.Split(i, "-")
				minPort, _ := strconv.Atoi(iSlice[0])
				maxPort, _ := strconv.Atoi(iSlice[1])

				for j := minPort; j <= maxPort; j++ {
					portSlice = append(portSlice, strconv.Itoa(j))
				}
			} else {

				portSlice = append(portSlice, i)
			}
		}
	}
	return portSlice
}

func parser() {
	Args = new(Options)
	flag.StringVar(&Args.ListenRange, "l", "normal", `监听范围，有三个选择
	常见81个端口: normal
	所有端口: all
	自定义：80-90,22`)
	flag.StringVar(&Args.HeaderFlag, "f", "test", `header flag
	内置: ldap/rmi
	自定义: flag或0x414141`)

	flag.Parse()

	if strings.HasPrefix(Args.HeaderFlag, "0x") {
		tmp, err := hex.DecodeString(Args.HeaderFlag[2:])
		if err != nil {
			log.LogError(err.Error())
			os.Exit(1)
		}
		Args.HeaderFlag = string(tmp)
	}
}
