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

func TestDiscoverLUNsPass(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.DiscoverLUNs("192.168.2.3")
	if err != nil {
		t.Error(err)
	}
}

func TestDiscoverLUNsFail(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.DiscoverLUNs("")
	if err == nil {
		t.Error("For no host - Functional error")
	}
}

func TestListVolumes(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.ListVolumes()
	if err != nil {
		t.Error(err)
	}
}

func TestLoginTargetAll(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.LoginTarget("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestLoginTargetOne(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.LoginTarget("abc.xyz.com", "10.0.2.15:3260")
	if err != nil {
		t.Error(err)
	}
}

func TestLogoutTargetAll(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.LogoutTarget("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestLogoutTargetOne(t *testing.T) {
	plugin := NewISCSIPlugin()
	err := plugin.LogoutTarget("abc.xyz.com", "10.0.2.15:3260")
	if err != nil {
		t.Error(err)
	}
}
