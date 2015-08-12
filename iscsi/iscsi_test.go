package iscsi

import "testing"

func TestExecuteCommand(t *testing.T) {
	out, err := ExecuteCommand("ls", "-l")
	t.Log(out)
	t.Log(err)
}

func TestCheckIscsiSupport(t *testing.T) {
	plugin := NewISCSIPlugin()
	isInstalled := plugin.CheckIscsiSupport()
	if !isInstalled {
		t.Error("iscsiadm not installed")
	}
}
