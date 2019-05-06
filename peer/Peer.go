package peer

import (
	"net"
)

type Peer struct {
	Ip         string
	Port       string
	ServerPort string
	Conn       net.Conn

	IsConnectedBack bool

	// Sending Queue?
}

func NewPeer(ip string, port string, conn net.Conn) *Peer {
	p := &Peer{ip, port, "6789", conn, false}
	return p
}

type PeerQueue struct {
	Peers []*Peer
}

func (pq *PeerQueue) EnQueue(peer *Peer) {
	pq.Peers = append(pq.Peers, peer)
}

func (pq *PeerQueue) GetUnSettlePeers() []*Peer {

	var result []*Peer

	for _, p := range pq.Peers {
		if p.IsConnectedBack == false {
			result = append(result, p)
		}
	}

	return result
}

func (pq *PeerQueue) GetSettledPeers() []*Peer {
	var result []*Peer

	for _, p := range pq.Peers {
		if p.IsConnectedBack == true {
			result = append(result, p)
		}
	}

	return result
}
