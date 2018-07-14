package utils

import (
	"fmt"
)

type ServerInfo struct {
	Host  string
	Port  int
	Proto string
}

func NewServer(host, proto string, port int) (*ServerInfo, error) {
	if len(host) == 0 || len(proto) == 0 || port <= 0 {
		return nil, fmt.Errorf("host,proto or port error.\n")
	}
	return &ServerInfo{
		Host:  host,
		Port:  port,
		Proto: proto,
	}, nil
}
