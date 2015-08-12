package iscsi

import (
	"bytes"
	_ "fmt"
	_ "log"
	"os/exec"
	"strings"
)

type ISCSIPlugin struct {
	hosts string
}

func ExecuteCommand(command string, args ...string) (string, string) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	var errMsg bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errMsg
	err := cmd.Run()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", "Command Not Found"
		}
		return "", errMsg.String()
	}
	return out.String(), errMsg.String()
}

func NewISCSIPlugin() ISCSIPlugin {
	return ISCSIPlugin{
		"test1",
	}
}

func (plugin *ISCSIPlugin) CheckIscsiSupport() bool {
	//Check if "iscsiadm" is installed
	_, err := ExecuteCommand("iscsiadm")
	if strings.Contains(err, "Command Not Found") {
		return false
	}
	return true
}

func (plugin *ISCSIPlugin) DiscoverLUNs(host string) error {
	return nil
}

func (plugin *ISCSIPlugin) ListVolumes() error {
	return nil
}

func (plugin *ISCSIPlugin) AddVolumes(volume string) error {
	return nil

}

func (plugin *ISCSIPlugin) DeleteVolumes(volume string) error {
	return nil

}
