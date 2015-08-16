package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/calavera/dkvolume"
	"github.com/codegangsta/cli"
	"github.com/phoenix-io/docker-iscsi-volume/iscsi"
)

const (
	iscsiConf     = "/etc/iscsi/iscsid.conf"
	iscsiVolumeID = "_iscsiVolume"
	socketAddress = "/usr/share/docker/plugins/iscsi-vol.sock"
)

var (
	defaultPath = filepath.Join(dkvolume.DefaultDockerRootDirectory, iscsiVolumeID)
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
			Name:  "discover",
			Usage: "Perform volume discovery",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target",
					Usage: "target IP / hostname for LUN discovery",
				},
			},
			Action: discoverVolumes,
		},
		{
			Name:   "login",
			Usage:  "login the target",
			Action: loginTarget,
		},
		{
			Name:   "logout",
			Usage:  "logout the target",
			Action: logoutTarget,
		},
	}
	plugin.Run(os.Args)

	d := newISCSIVolumeDriver("iscsi")
	h := dkvolume.NewHandler(d)
	fmt.Println("Listening on ", socketAddress)
	fmt.Println(h.ServeUnix("root", socketAddress))

}

func listVolumes(c *cli.Context) {
	plugin := iscsi.NewISCSIPlugin()
	err := plugin.ListVolumes()
	if err != nil {
		log.Panic(err)
	}
}

func discoverVolumes(c *cli.Context) {
	target := c.String("target")
	if len(target) == 0 {
		cli.ShowCommandHelp(c, "discover")
		return
	}

	plugin := iscsi.NewISCSIPlugin()
	err := plugin.DiscoverLUNs(target)
	if err != nil {
		log.Panic(err)
	}
}

func loginTarget(c *cli.Context) {
	plugin := iscsi.NewISCSIPlugin()
	err := plugin.LoginTarget("", "")
	if err != nil {
		log.Panic(err)
	}
}

func logoutTarget(c *cli.Context) {
	plugin := iscsi.NewISCSIPlugin()
	err := plugin.LogoutTarget("", "")
	if err != nil {
		log.Panic(err)
	}
}
