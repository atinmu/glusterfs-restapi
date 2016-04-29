package main

import (
	"encoding/json"
	"net/http"

	"gluster/cli"
	"gluster/utils"
)

// PeersAdd is a Handler function to add nodes to Gluster Cluster
func PeersAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var peerNodes []string
	err := decoder.Decode(&peerNodes)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, peer := range peerNodes {
		errPeerAdd := cli.PeerAttach(peer)
		if errPeerAdd != nil {
			utils.HTTPErrorJSON(w, errPeerAdd.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// PeersRemove is a Handler function to remove nodes to Gluster Cluster
func PeersRemove(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var peerNodes []string
	err := decoder.Decode(&peerNodes)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, peer := range peerNodes {
		errPeerRemove := cli.PeerDitach(peer)
		if errPeerRemove != nil {
			utils.HTTPErrorJSON(w, errPeerRemove.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// PeersGet is a Handler func to get list of peers of Gluster Cluster
func PeersGet(w http.ResponseWriter, r *http.Request) {
	var info []cli.Peer
	var err error
	info, err = cli.PoolList()
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	utils.HTTPOutJSON(w, info)
}
