package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// Config to store all configurations related to REST
type Config struct {
	AuthEnabled     bool          `json:"auth_enabled"`
	Port            int           `json:"port"`
	UseHTTPS        bool          `json:"https"`
	Csr             string        `json:"csr"`
	Key             string        `json:"key"`
	AppsFile        string        `json:"apps_file"`
	AccessLogFile   string        `json:"access_log_file"`
	EventsSockFile  string        `json:"events_sock_file"`
	InternalUser    string        `json:"internal_user"`
	ListenURL       string        `json:"listen_url"`
	APIVersion      string        `json:"api_version"`
	EventsURL       string        `json:"events_url"`
	WebsocketExpiry time.Duration `json:"websocket_expiry"`
}

func loadConfig(defaultConfigFile string, customConfigFile string, fail bool) {
	data, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil && fail {
		Logger.Fatal("No conf file")
	}

	err1 := json.Unmarshal(data, &RestConfig)
	if err1 != nil && fail {
		log.Fatal("json err")
	}

	// Reading Custom Config File
	data1, err2 := ioutil.ReadFile(customConfigFile)
	if err2 != nil {
		log.Println("No custom conf file")
		return
	}

	err3 := json.Unmarshal(data1, &RestConfig)
	if err3 != nil && fail {
		log.Fatal("json err custom config file")
	}
}
