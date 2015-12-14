package comm

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Create(isType string) {

	if isType == "server" {
		server()
	} else {
		client()
	}
}

func server() {

	fmt.Println("Run server...")
	listen, _ := net.Listen("tcp", ":8081")

	conn, _ := listen.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Received message: ", string(msg))

		newMsg := strings.ToUpper(msg)
		conn.Write([]byte(newMsg + "\n"))
	}
}

func client() {

	fmt.Println("Run client...")
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send message: ")

		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message form server: " + message)
	}
}
