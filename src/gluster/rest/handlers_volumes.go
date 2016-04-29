package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gluster/cli"
	"gluster/utils"
)

// VolumeCreate is a Handler function to create Gluster Volume
func VolumeCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var opts cli.CreateOptions
	err := decoder.Decode(&opts)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	volName := vars["volName"]
	errCreate := cli.VolumeCreate(volName, opts.Bricks, opts)
	if errCreate != nil {
		utils.HTTPErrorJSON(w, errCreate.Error(), http.StatusInternalServerError)
		return
	}
}

// VolumeGet is a HTTP Handler function to get Gluster Volume Information
func VolumeGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName, ok := vars["volName"]
	if !ok {
		volName = ""
	}
	status := r.URL.Query().Get("status")
	var info []cli.Volume
	var err error
	if status == "1" {
		info, err = cli.VolumeStatus(volName)
	} else {
		info, err = cli.VolumeInfo(volName)
	}
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	utils.HTTPOutJSON(w, info)
}

// VolumeStatus is a HTTP Handler function to get Gluster Volume Status
func VolumeStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName, ok := vars["volName"]
	if !ok {
		volName = ""
	}

	info, err := cli.VolumeStatus(volName)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	utils.HTTPOutJSON(w, info)
}

// VolumeStart is a HTTP handler to Start Gluster Volume
func VolumeStart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := cli.VolumeStart(volName, false)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// VolumeStop is a HTTP handler to Stop Gluster Volume
func VolumeStop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := cli.VolumeStop(volName, false)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// VolumeDelete is a HTTP handler to Delete a Gluster Volume
func VolumeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := cli.VolumeDelete(volName)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// VolumeOptionsGet is a HTTP handler func to
func VolumeOptionsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName, ok := vars["volName"]
	if !ok {
		utils.HTTPErrorJSON(w, "Invalid Volume name", http.StatusBadRequest)
		return
	}
	all := r.URL.Query().Get("all")
	var info []cli.VolumeOption
	var err error
	if all == "1" {
		info, err = cli.VolumeOptGet(volName, "all")
	} else {
		info, err = cli.VolumeOptGet(volName, "")
	}

	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	utils.HTTPOutJSON(w, info)
}

func VolumeOptionsSet(w http.ResponseWriter, r *http.Request) {
	var opts = make(map[string]string)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&opts)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	volName := vars["volName"]
	for k, v := range opts {
		err := cli.VolumeOptSet(volName, k, v)
		if err != nil {
			utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func VolumeOptionsReset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]

	all := r.URL.Query().Get("all")
	if all == "1" {
		err := cli.VolumeOptReset(volName, "", false)
		if err != nil {
			utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Reset specific keys
	var opts []string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&opts)
	if err != nil {
		utils.HTTPErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, k := range opts {
		err := cli.VolumeOptReset(volName, k, false)
		if err != nil {
			utils.HTTPErrorJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
