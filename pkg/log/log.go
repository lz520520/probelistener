package log

import "fmt"

func LogInfo(msg string) {
	fmt.Println("[+] " + msg)
}

func LogError(msg string) {
	fmt.Println("[-] " + msg)
}
