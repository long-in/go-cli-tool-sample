package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cli.Command{
			Name: "add",
			//Aliases: []string{"a"},
			Usage: "`add` is an example of an option.",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "flag, f"},
			},
			Action: func(c *cli.Context) error {
				flag := c.String("flag")
				if flag == "" {
					fmt.Println("--flag a option not found.")
					return nil
				}
				return nil
			},
		},
		cli.Command{
			Name: "del",
			//Aliases: []string{"d"},
			Usage: "`add` is an example of an option.",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "flag, f"},
				cli.StringFlag{Name: "flag2, f2"},
			},
			Action: func(c *cli.Context) error {
				flag := c.String("flag")
				flag2 := c.String("flag2")
				if flag == "" || flag2 == "" {
					fmt.Println("--flag or --flag2 option not found.")
					return nil
				}
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	check(err)
}
