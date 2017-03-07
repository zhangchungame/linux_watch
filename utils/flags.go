package utils

import (
	"fmt"
	"github.com/urfave/cli"
)

func NewApp()  *cli.App{
	app := cli.NewApp()
	app.Name = "linux_watch"
	app.Usage = "watch linux"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	return app
}

