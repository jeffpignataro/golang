package main

import (
	"github.com/urfave/cli/v2"
)

func Commands(c *cli.Context) []*cli.Command {
	return []*cli.Command{
		firstCommand(c),
	}
}

func firstCommand(c *cli.Context) *cli.Command {
	return &cli.Command{
		Name:    "FirstCommand",
		Aliases: []string{"f", "first"},
		Usage:   "Run an abstracted command.",
		Action:  firstAction(c),
		Subcommands: []*cli.Command{
			subCommand(c),
		},
	}
}

func subCommand(c *cli.Context) *cli.Command {
	return &cli.Command{
		Name:        "SubCommand",
		Aliases:     []string{"sub"},
		Usage:       "Run sub command of first command",
		Action:      subAction(c),
		Subcommands: nil,
	}
}
