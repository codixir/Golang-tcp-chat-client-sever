package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	color "github.com/fatih/color"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)

	defer connection.Close()

	color.Cyan("Enter your username:")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = username[:len(username)-1]

	welcomeMsg := fmt.Sprintf("Welcome %s, say hi to your friends:\n", username)
	color.Magenta(welcomeMsg)

	go read(connection)
	write(connection, username)
}

func read(connection net.Conn) {

	reader := bufio.NewReader(connection)
	x, _ := reader.ReadString('\n')

	color.Yellow(x)

	for {
		reader = bufio.NewReader(connection)
		message, err := reader.ReadString('\n')

		if err == io.EOF {
			connection.Close()
			color.Red("Connection closed, see you later!")
			os.Exit(0)
		}

		color.Green(message)
		color.Magenta("------------------------------------")
	}
}

func write(connection net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = fmt.Sprintf("%s: %s\n", username,
			message[:len(message)-1])

		connection.Write([]byte(message))
	}
}
