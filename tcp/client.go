package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"programacao-em-rede/util"
	"strconv"
	"strings"
)

func main() {

	ip := util.SERVER_IP
	port := strconv.Itoa(util.SERVER_PORT)
	endpoint := ip + ":" + port

	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Client >> ")
		request, _ := reader.ReadString('\n')
		fmt.Fprint(conn, request+"\n")

		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server >> " + response)

		if strings.TrimSpace(string(request)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
