package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	EOS  = "\000\111\222\333"
	PORT = ":8080"
)

func InputString() string {
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(msg, "\n", "", -1)
}

func clientInput(conn net.Conn) {
	for {
		conn.Write([]byte(InputString() + EOS))
	}
}

func clientOutput(conn net.Conn) {
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
			fmt.Println(messgae)
		}
		fmt.Println(messgae)
	}
}

func main() {
	conn, err := net.Dial("tcp", PORT)
	if err != nil {
		panic("cant connect to server")
	}
	defer conn.Close()

	go clientOutput(conn)
	clientInput(conn)
}
