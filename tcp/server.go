package main

import (
	"bufio"
	"fmt"
	"net"
	"programacao-em-rede/util"
	"strconv"
	"strings"
)

func main() {

	port := strconv.Itoa(util.SERVER_PORT)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	fmt.Println("Server Listening...")

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(conn)
		return
	}

	for {
		request, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(request)) == "STOP" {
			fmt.Println("Exiting TCP Server")
			return
		}

		fmt.Print("-> ", string(request))
		response := strings.ToUpper(request) + "\n"
		conn.Write([]byte(response))
	}
}
