package main

import (
	"github.com/vinsia/fly/fly"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func initFlags() (app *cli.App) {
	app = cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "mode, m",
			Usage: "mode",
		},
		cli.StringFlag{
			Name:  "name, n",
			Usage: "server name",
		},
		cli.StringFlag{
			Name:  "host",
			Usage: "server host",
		},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "user name",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "login password",
		},
	}
	return
}

func main() {
	app := initFlags()
	app.Action = func(context *cli.Context) (err error) {
		_fly := fly.NewFly()
		answer := _fly.Ask()
		log.Println("123123",answer)
		log.Println(_fly)
		a := _fly.Run(answer)
		log.Println(a)
		log.Println("444")
		return
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("can not parse arguments: %e", err)
	}
}
