package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gluster/utils"
)

func main() {
	utils.LogInit(logFilePath)
	utils.Autoload(defaultConfigPath, customConfigPath)
	router := mux.NewRouter().StrictSlash(true)
	AddRoutes(router)

	portData := fmt.Sprintf(":%d", utils.RestConfig.Port)
	utils.Logger.Info("Started running REST server in port ", utils.RestConfig.Port)
	if utils.RestConfig.UseHTTPS {
		utils.Logger.Fatal(http.ListenAndServeTLS(portData, utils.RestConfig.Csr, utils.RestConfig.Key, nil))
	} else {
		utils.Logger.Fatal(http.ListenAndServe(portData, nil))
	}
}
