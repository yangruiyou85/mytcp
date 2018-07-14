package tcpclient

import (
	"net"
	"bufio"
	"fmt"
	"log"
	"mytcp/cmd/pkg/utils"
)

func ConnectServer(s *utils.ServerInfo) {

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	conn, err := net.Dial(s.Proto,addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Fprintf(conn, "this is client...")
	content, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)

}
