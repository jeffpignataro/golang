package main

import (
	"fmt"

	doathing "github.com/jeffpignataro/golang/pkg/do-a-thing"
	doathingv2 "github.com/jeffpignataro/golang/pkg/do-a-thing/v2"
	"github.com/urfave/cli/v2"
)

func firstAction(c *cli.Context) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Println(doathing.Doathing(c.Args().First()))
		fmt.Println(doathingv2.Doathing(c.Args().First()))
		return nil
	}
}

func subAction(c *cli.Context) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Println("This is the first commands sub command.")
		return nil
	}
}
