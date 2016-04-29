package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
)

var (
	// RestConfig will have all the configurations once Autoload is executed.
	RestConfig Config
	// PeersList will have peers list once Autoload is executed.
	PeersList Peers
	// RestApps will have peers list once Autoload is executed.
	RestApps = make(Apps)
	// MyUUID is local node UUID loaded after Autoload is executed.
	MyUUID = ""
	// Logger is instance created for Logging
	Logger = logrus.New()
)

// CmdResponse is used to return the output of Gluster Command execution.
type CmdResponse struct {
	Ok         bool   `json:"ok"`
	Msg        string `json:"msg"`
	ReturnCode string `json:"return_code"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// HTTPErrorJSON is a utility function to write error in JSON format
// to HTTP ResponseWriter
func HTTPErrorJSON(w http.ResponseWriter, err string, code int) {
	msg := errorResponse{Message: err}
	j, _ := json.Marshal(msg)
	w.WriteHeader(code)
	w.Write(j)
}

// HTTPOutJSON is a utility func to write JSON output to given HTTP
// ResponseWriter
func HTTPOutJSON(w http.ResponseWriter, out interface{}) {
	j, _ := json.Marshal(out)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// LogInit is a utility function to initialize logging
func LogInit(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Unable to open Log file ", err)
	}
	Logger.Out = f
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2016-03-28 15:04:05"
	Logger.Formatter = customFormatter
	customFormatter.FullTimestamp = true
	customFormatter.DisableColors = true
}

// Autoload is a utility function to load config, Apps and Peers list
// It reloads the config, apps and peers list When it recieve SIGUSR2.
func Autoload(defaultConfigFile string, customConfigFile string) {
	loadConfig(defaultConfigFile, customConfigFile, true)
	loadApps(true)
	loadPeers(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			loadConfig(defaultConfigFile, customConfigFile, false)
			loadApps(false)
			loadPeers(false)
			Logger.Println("Reloaded config, apps and Peers List")
		}
	}()
}

// Execute is a helper func to execute Gluster Commands
func Execute(cmd []string) CmdResponse {
	cmd = append([]string{"--mode=script"}, cmd...)
	out := CmdResponse{Ok: true}
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		out.Ok = false
		out.Msg = strings.Trim(string(o), "\n")
	}
	return out
}

// GetQsh will generate qsh by using Method, Path and Data.
func GetQsh(method string, path string, queryParams string, data string) string {
	qshData := method + "\n" + path
	if queryParams != "" {
		qshData = qshData + "\n" + queryParams
	}

	if data != "" {
		qshData = qshData + "\n" + data
	}
	hash := sha256.New()
	hash.Write([]byte(qshData))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sign is a utility func to generate JWT token using the secret and
// Claims
func Sign(secret string, iss string, qsh string) string {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["iss"] = iss
	if qsh != "" {
		token.Claims["qsh"] = qsh
	}
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenString
}
