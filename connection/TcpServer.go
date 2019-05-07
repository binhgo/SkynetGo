package connection

import (
	"fmt"
	"log"
	"net"

	"SkynetGo/peer"
)

func InitServer(TcpPort string, queue *peer.PeerQueue) *net.Conn {

	fmt.Printf("Server is listening on port %s\n", TcpPort)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", TcpPort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()

		log.Printf("Accepted conn from %s\n", conn.RemoteAddr().String())

		if err != nil {
			continue
		}

		ip, port := getIPAndPort(conn.RemoteAddr())
		p := peer.NewPeer(ip, port, conn)
		queue.Peers = append(queue.Peers, p)

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 1024)

	for {
		len, err := conn.Read(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len != 0 {
			fmt.Printf("Receive data from client: %s\n", string(data[:len]))
		}
	}
}

func SendDataToClient(conn net.Conn, data string) {
	if conn != nil {
		_, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func SendDataToAllClients(peerQueue *peer.PeerQueue, data string) {
	for _, p := range peerQueue.GetSettledPeers() {
		SendDataToClient(p.Conn, data)
	}
}

func RegisterWithGlobalPeerQueue() {

}

func GetGlocalPeerQueue() {

}
