package main

import (
	"bufio"
	"fmt"
	"net"
	"programacao-em-rede/util"
	"strconv"
	"strings"
)

var count = 0

func main() {

	port := strconv.Itoa(util.SERVER_PORT)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	fmt.Println("Server Listening...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(conn)
			return
		}
		go newConnection(conn)
		count++
	}
}

func newConnection(conn net.Conn) {

	for {
		request, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(request)) == "STOP" {
			fmt.Println("Exiting TCP Server")
			break
			return
		}

		fmt.Print("-> ", string(request))
		counter := strconv.Itoa(count)
		fmt.Print(counter)
		response := strings.ToUpper(request) + " Client: " + counter + "\n"
		conn.Write([]byte(response))
	}
}
