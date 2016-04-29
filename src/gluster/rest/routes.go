package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	// Volume Life Cycle APIs
	router.HandleFunc("/v1/volumes/{volName}", VolumeCreate).Methods("PUT")
	router.HandleFunc("/v1/volumes/{volName}/start", VolumeStart).Methods("POST")
	router.HandleFunc("/v1/volumes/{volName}/stop", VolumeStop).Methods("POST")
	router.HandleFunc("/v1/volumes/{volName}", VolumeDelete).Methods("DELETE")
	router.HandleFunc("/v1/volumes/{volName}", VolumeGet).Methods("GET")
	router.HandleFunc("/v1/volumes", VolumeGet).Methods("GET")

	// Volume Options
	router.HandleFunc("/v1/volumes/{volName}/options", VolumeOptionsGet).Methods("GET")
	router.HandleFunc("/v1/volumes/{volName}/options", VolumeOptionsSet).Methods("POST")
	router.HandleFunc("/v1/volumes/{volName}/options", VolumeOptionsReset).Methods("DELETE")

	// Peers
	router.HandleFunc("/v1/peers", PeersAdd).Methods("POST")
	router.HandleFunc("/v1/peers", PeersRemove).Methods("DELETE")
	router.HandleFunc("/v1/peers", PeersGet).Methods("GET")

	http.Handle("/",
		RestLoggingHandler(
			SetApplicationHeaderJSON(
				VerifyHandler(router))))
}
