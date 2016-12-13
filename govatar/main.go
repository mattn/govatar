package main

import (
	"errors"
	"os"

	"github.com/o1egl/govatar"
	"github.com/urfave/cli"
)

var version = "0.1.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "govatar"
	app.Usage = "Avatar generator service."
	app.Version = version
	app.Author = "Oleg Lobanov"
	app.Commands = []cli.Command{
		{
			Name:      "generate",
			ArgsUsage: "<(male|m)|(female|f)>",
			Aliases:   []string{"g"},
			Usage:     "Generates random avatar",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output,o",
					Value: "avatar.png",
					Usage: "Output file name",
				},
			},
			Action: func(c *cli.Context) error {
				var g govatar.Gender
				switch c.Args().First() {
				case "male", "m":
					g = govatar.MALE
				case "female", "f":
					g = govatar.FEMALE
				default:
					return errors.New("Incorrect gender param. Run `govatar help generate`")
				}
				err := govatar.GenerateFile(g, c.String("output"))
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
