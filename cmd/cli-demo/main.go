package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var flag1Value string

func main() {
	app := &cli.App{
		Name:           "demo",
		HelpName:       "",
		Usage:          "A sample CLI application",
		UsageText:      "",
		ArgsUsage:      "",
		Version:        "",
		Description:    "",
		DefaultCommand: "",
		Commands:       Commands(&cli.Context{}),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "flag1",
				Usage:       "This is a demo flag",
				Destination: &flag1Value,
			},
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(*cli.Context) {
		},
		Before: func(*cli.Context) error {
			return nil
		},
		After: func(*cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			if flag1Value == "" {
				fmt.Printf("Hello %s", c.Args().Get(0))
			} else {
				fmt.Printf("This is a flag execution: %s", flag1Value)
			}
			return nil
		},
		CommandNotFound: func(*cli.Context, string) {
		},
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		InvalidFlagAccessHandler: func(*cli.Context, string) {
		},
		Compiled:  time.Time{},
		Authors:   []*cli.Author{},
		Copyright: "",
		Reader:    nil,
		Writer:    nil,
		ErrWriter: nil,
		ExitErrHandler: func(cCtx *cli.Context, err error) {
		},
		Metadata: map[string]interface{}{},
		ExtraInfo: func() map[string]string {
			return nil
		},
		CustomAppHelpTemplate:  "",
		SliceFlagSeparator:     "",
		UseShortOptionHandling: false,
		Suggest:                false,
		AllowExtFlags:          false,
		SkipFlagParsing:        false,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
