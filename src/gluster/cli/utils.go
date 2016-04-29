package cli

import (
	"errors"
	"os/exec"
	"strings"
)

// ExecuteCmd is helper function to execute Gluster Command
func ExecuteCmd(cmd []string) error {
	cmd = append([]string{"--mode=script"}, cmd...)
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		return errors.New(strings.Trim(string(o), "\n"))
	}
	return nil
}

// ExecuteCmdXML is helper function to execute Gluster Command with `--xml` option
func ExecuteCmdXML(cmd []string) ([]byte, error) {
	cmd = append([]string{"--mode=script", "--xml"}, cmd...)
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		return []byte(""), errors.New(strings.Trim(string(o), "\n"))
	}
	return o, nil
}
