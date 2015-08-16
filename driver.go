package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/calavera/dkvolume"
)

type iscsiDriver struct {
	baseMountPath string
	m             *sync.Mutex
}

func newISCSIVolumeDriver(root string) iscsiDriver {
	return iscsiDriver{
		baseMountPath: root,
		m:             &sync.Mutex{},
	}
}

func (d iscsiDriver) Create(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

func (d iscsiDriver) Remove(r dkvolume.Request) dkvolume.Response {
	d.m.Lock()
	defer d.m.Unlock()

	_ = d.mountpoint(r.Name)

	// logout all mountpoints

	return dkvolume.Response{}
}

func (d iscsiDriver) Path(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{Mountpoint: d.mountpoint(r.Name)}
}

func (d iscsiDriver) Mount(r dkvolume.Request) dkvolume.Response {
	d.m.Lock()
	defer d.m.Unlock()
	m := d.mountpoint(r.Name)
	log.Printf("Mounting volume %s on %s\n", r.Name, m)

	//Create a temp folder in mountPath.
	os.Mkdir(m, os.ModeDir)
	plugin := NewISCSIPlugin()
	err := plugin.LoginTarget(plugin., "10.0.2.15:3260")
	if err != nil {
		t.Error(err)
	}
	//Login te target
	//Mount logic.
	return dkvolume.Response{Mountpoint: m}
}

func (d iscsiDriver) Unmount(r dkvolume.Request) dkvolume.Response {
	d.m.Lock()
	defer d.m.Unlock()
	m := d.mountpoint(r.Name)
	log.Printf("Unmount volume %s on %s\n", r.Name, m)

	//Create a temp folder in mountPath.
	//Login te target
	//Mount logic.
	return dkvolume.Response{Mountpoint: m}
}

func (d iscsiDriver) mountpoint(name string) string {
	return filepath.Join(d.baseMountPath, name)
}
