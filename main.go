package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"go.universe.tf/netboot/tftp"
)

var BuildVersion string = "0.1.0"

func terminalPath() string {
	path, _ := os.Getwd()
	return path
}

func main() {
	app := cli.NewApp()
	app.Version = BuildVersion
	app.Description = "CLI TFTP server"
	app.DefaultCommand = "run"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Value: terminalPath(), Usage: "TFTP service root path (supports both absolute and relative paths)", DefaultText: "relative path"},
		&cli.StringFlag{Name: "listen", Aliases: []string{"l"}, Value: "", Usage: "TFTP service listening address", DefaultText: ":69"},
	}
	app.Commands = []*cli.Command{
		{
			Name: "run", Action: func(ctx *cli.Context) error {
				path := ctx.String("path")
				if !filepath.IsAbs(path) {
					path = filepath.Join(terminalPath(), path)
				}
				if hnd, err := tftp.FilesystemHandler(path); err == nil {
					fmt.Println("ROOT", path)
					ser := &tftp.Server{Handler: hnd}
					return ser.ListenAndServe(ctx.String("listen"))
				} else {
					return err
				}
			},
			Usage: "Default command. Start the TFTP service.",
		},
	}
	if err := app.Run(os.Args); err != nil {
		return
	}
}
