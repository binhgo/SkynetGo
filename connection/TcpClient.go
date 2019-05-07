package connection

import (
	"fmt"
	"net"

	"SkynetGo/peer"
)

func InitClient(peer *peer.Peer) {

	IPAndPort := fmt.Sprintf("%s:%s", peer.Ip, peer.ServerPort)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", IPAndPort)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// set peer is settled
	peer.IsConnectedBack = true

	go receiveData(conn)
}

// Send data to server
func SendDataToServer(conn net.Conn, data string) {
	fmt.Printf("%s", "Client send data...\n")

	_, err := conn.Write([]byte(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}

func receiveData(conn net.Conn) {
	fmt.Printf("%s%s\n", "Client start receiving data...", conn)
	data := make([]byte, 1024)

	for {
		len, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len != 0 {
			fmt.Printf("Client receive data: %s from %s\n", string(data[0:len]), conn.RemoteAddr().String())
		}
	}
}
