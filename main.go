package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"SkynetGo/Connection"
	"SkynetGo/peer"
)

var TCPPORT = ":6789"
var peerQueue *peer.PeerQueue = new(peer.PeerQueue)

func main() {
	// Mock
	tempPeers := CreatePeerQueue()
	//go LogPeerQueue()

	// Init Server at first to listen to all incoming messages
	go Connection.InitServer(TCPPORT, peerQueue)
	// Mock
	go InitFAKEClients(tempPeers)
	go UpdateFAKEClients()

	// go InitClients()

	go TestSendToClient()

	waitExitSignal()
}

func InitFAKEClients(clients []*peer.Peer) {
	time.Sleep(time.Second * 2)

	// Check peers list and setup new clients connect to these peer's server
	for _, p := range clients {
		Connection.InitClient(p)
	}

}

func UpdateFAKEClients() {

	time.Sleep(time.Second * 4)

	for _, p := range peerQueue.Peers {
		p.IsConnectedBack = true
	}
}

func InitClients() {
	for {
		time.Sleep(time.Second * 20)
		peers := peerQueue.GetUnSettlePeers()

		// Check peers list and setup new clients connect to these peer's server
		for _, p := range peers {
			Connection.InitClient(p)
		}
	}
}

func TestSendToClient() {
	count := 0
	for {
		time.Sleep(time.Second * 6)

		count++
		message := fmt.Sprintf("Hello from SERVER. Count %d", count)
		Connection.SendDataToAllClients(peerQueue, message)
	}
}

func CreatePeerQueue() []*peer.Peer {

	var peers []*peer.Peer

	p1 := peer.NewPeer("localhost", "", nil)
	peers = append(peers, p1)

	p2 := peer.NewPeer("localhost", "", nil)
	peers = append(peers, p2)

	return peers
}

func LogPeerQueue() {
	for {
		time.Sleep(time.Second * 3)
		for _, p := range peerQueue.Peers {
			log.Printf("Queue: %s", p)
		}
	}
}

func waitExitSignal() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	<-done
}
