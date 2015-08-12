package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/calavera/dkvolume"
)

type iscsiDriver struct {
	m             *sync.Mutex
	baseMountPath string
}

func NewISCSIVolumeDriver(root string) {
	return iscsiDriver{
		m:             &sync.Mutex{},
		baseMountPath: root,
	}
}

func (d *iscsiDriver) Create(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

func (d *iscsiDriver) Remove(r dkvolume.Request) dkvolume.Response {
	d.m.Lock()
	defer d.m.Unlock()

	m := d.mountpoint(r.Name)

	// logout all mountpoints

	return dkvolume.Response{}
}

func (d *iscsiDriver) Path(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{Mountpoint: d.mountpoint(r.Name)}
}

func (d *iscsiDriver) Mount(r dkvolume.Request) dkvolume.Response {
	d.m.Lock()
	defer d.m.Unlock()
	m := d.mountpoint(r.Name)
	log.Printf("Mounting volume %s on %s\n", r.Name, m)

	//Create a temp folder in mountPath.
	//Login te target
	//Mount logic.
	return dkvolume.Response{Mountpoint: m}
}

func (d *iscsiDriver) mountpoint(name string) string {
	return filepath.Join(d.baseMountPath, name)
}
