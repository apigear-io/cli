package helper

import (
	"log"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	if err := l.Close(); err != nil {
		log.Printf("error closing listener: %v", err)
		_ = err
	}
	return l.Addr().(*net.TCPAddr).Port, nil
}
