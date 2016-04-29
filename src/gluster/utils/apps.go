package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Apps to store the applications ID:Secret
type Apps map[string]string

func loadApps(fail bool) {
	data, err := ioutil.ReadFile(RestConfig.AppsFile)
	if err != nil && fail {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal("No apps file")
	}

	err1 := json.Unmarshal(data, &RestApps)
	if err1 != nil && fail {
		log.Fatal("json err")
	}
}
