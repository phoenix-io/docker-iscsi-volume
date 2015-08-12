package main

import (
	"fmt"
	_ "log"
	"os"

	_"github.com/calavera/dkvolume"
	"github.com/codegangsta/cli"
)

const (
	iscsiConf      = "/etc/iscsi/iscsid.conf"
	socketAddress = "/usr/share/docker/plugins/iscsi-vol.sock"
)


func main() {

	plugin := cli.NewApp()
	plugin.Name = "iscsi-docker-plugin"
	plugin.Usage = "Manage iSCSI Volumes"
	plugin.Version = "0.1.0"
	plugin.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "List the iSCSI volumes (added/discovered)",
			Action: listVolumes,
		},
		{
			Name:   "discover",
			Usage:  "Perform volume discovery",
			Action: discoverVolumes,
		},
		{
			Name:   "add",
			Usage:  "Adds the volume",
			Action: addVolume,
		},
		{
			Name:   "del",
			Usage:  "Deletes the volume",
			Action: delVolume,
		},
	}
	plugin.Run(os.Args)
	// Create Plugin Driver
	// SetNew Handler.
	// Listen at socket.
}

func listVolumes(c *cli.Context) {
	fmt.Println("NOT IMPLEMENTED")
}

func discoverVolumes(c *cli.Context) {
	fmt.Println("NOT IMPLEMENTED")
}

func addVolume(c *cli.Context) {
	fmt.Println("NOT IMPLEMENTED")
}

func delVolume(c *cli.Context) {
	fmt.Println("NOT IMPLEMENTED")
}
