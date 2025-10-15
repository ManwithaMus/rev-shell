package main

import(
		"bufio"
		"fmt"
		"log"
		"net"
)

func main(){
	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalf("Error listening: %v", err) // Logs error when listening fails
	}
	defer listener.Close()

	fmt.Println("Listening for incoming connections on port 4444...")

	for {

		connection, err := listener.Accept()
		if err != nil{
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go connect(connection) // Establish an incoming connection
	}
}

func connect(conn net.Conn){
	defer conn.Close() // Close the connection at the end
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for{
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr())
			}else{
				log.Printf("Error reading from connection %s: %v", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Printf("Received from %s: %s", conn.RemoteAddr(), message) // Print response
	
		_, err = conn.Write([]byte("Echo: " + message))
		if err != nil {
			log.Printf("Error writing to connection %s: %v", conn.RemoteAddr(), err)
			return
		}
	}

}