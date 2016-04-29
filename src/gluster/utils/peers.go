package utils

import (
	"log"

	"gluster/cli"
)

// Peers to store Gluster Peers list
type Peers []cli.Peer

func loadPeers(fail bool) {
	p, err := cli.PoolList()
	if err != nil {
		if fail {
			log.Fatal("Peers list failed", err)
		}
	}
	PeersList = p

	for _, peer := range PeersList {
		if peer.Hostname == "localhost" {
			MyUUID = peer.ID
			break
		}
	}
}
