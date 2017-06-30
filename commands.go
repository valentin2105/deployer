package main

import (
	"fmt"
	"os"

	"git.nautile.tech/publicrepo/deployer/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "deploy",
		Usage:  "",
		Action: command.CmdDeploy,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "delete",
		Usage:  "",
		Action: command.CmdDelete,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
