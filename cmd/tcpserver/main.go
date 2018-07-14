package tcpserver

import (
	"fmt"
	"net"
	"github.com/gpmgo/gopm/modules/log"
	"mytcp/cmd/pkg/utils"
	"io"
)

func TcpServerCreate(s *utils.ServerInfo) {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	fmt.Println("Listening ......")
	l, err := net.Listen(s.Proto, addr)

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	for {
		fmt.Println("Accept......")
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
