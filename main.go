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
		_fly.Fly()
		defer _fly.Crash()
		if context.String("mode") != "" {
			if context.String("mode") == "add" {
				_fly.UpdateServer(fly.Server{
					Name:     context.String("name"),
					UserName: context.String("user"),
					Host:     context.String("host"),
					Password: context.String("password"),
				})
			} else if context.String("mode") == "default" {
				_fly.UpdateDefault(fly.Server{
					Name:     context.String("name"),
					UserName: context.String("user"),
					Host:     context.String("host"),
					Password: context.String("password"),
				})
			} else {
				log.Fatal("Not support such mode.")
			}
		} else {
			answer := _fly.Ask()
			_fly.Run(answer)
		}
		return
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("can not parse arguments: %e", err)
	}
}
