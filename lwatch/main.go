package main

import (
	"linux_watch/utils"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app:=utils.NewApp()
	app.Action=utils.AllWatch
	app.Commands=[]cli.Command{
		utils.Cpucommand,
	}
	app.Flags=[]cli.Flag{
		cli.StringFlag{Name: "yamlfile", Value: "", Usage: "配置文件位置",Destination:&utils.Yamlfile},
	}
	app.Run(os.Args)
}