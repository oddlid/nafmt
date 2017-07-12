// nafmt - nagios format
// Simple util that reads Nagios objects from stdin, then prints them formatted and sorted on stdout
package main

import (
	"github.com/urfave/cli"
	"github.com/vgtmnm/nagioscfg"
	"os"
)

const VERSION string = "2017-07-12"

func entryPoint(ctx *cli.Context) error {
	outfile := ctx.String("output")

	ncfg := nagioscfg.NewNagiosCfg()
	err := ncfg.LoadStdin()
	if err != nil {
		return cli.NewExitError("Unable to load STDIN", 1)
	}

	if outfile == "" {
		ncfg.Print(os.Stdout, true)
	} else {
		err := ncfg.WriteFile(outfile, true)
		if err != nil {
			return cli.NewExitError(err.Error(), 73) // EX_CANTCREAT=73 # can't create (user) output file
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Nagios Format"
	app.Version = VERSION
	app.Usage = "Parse and format Nagios objects in a pipe"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Odd E. Ebbesen",
			Email: "odd.ebbesen@wirelesscar.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Save output to `FILE` instead of STDOUT",
		},
	}

	app.Action = entryPoint
	app.Run(os.Args)
}
