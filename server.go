package main

import (
	"log"
	"net"
	"strings"
)

const (
	EOS  = "\000\111\222\333"
	PORT = ":8080"
)

var (
	Connections = make(map[net.Conn]bool)
)

func HandleConnect(conn net.Conn) {
	Connections[conn] = true
	log.Println(conn.RemoteAddr())
	var (
		messgae string
		buff    = make([]byte, 512)
	)
close:
	for {
		messgae = ""
		for {
			len, err := conn.Read(buff)
			if len == 0 || err != nil {
				break close
			}
			messgae += string(buff[:len])
			if strings.HasSuffix(messgae, EOS) {
				messgae = strings.TrimSuffix(messgae, EOS)
				break
			}
		}
		//conn.Write([]byte(strings.ToLower(messgae)))
		log.Println(messgae)
		for addr := range Connections {
			if addr == conn {
				continue
			}
			addr.Write([]byte(strings.ToLower(messgae) + EOS))
		}
	}
	/*for {
		len, err := conn.Read(buff)
		if len == 0 || err != nil {
			break
		}
		messgae += string(buff[:len])
		if strings.HasSuffix(messgae, EOS) {
			messgae = strings.TrimSuffix(messgae, EOS)
			break
		}
	}
	//conn.Write([]byte(strings.ToLower(messgae)))
	for addr := range Connections {
		if addr == conn {
			continue
		}
		addr.Write([]byte(strings.ToLower(messgae) + EOS))
	}
	log.Println(messgae)*/
	delete(Connections, conn)
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic("server not started")
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go HandleConnect(conn)
	}
}
