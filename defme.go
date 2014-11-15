package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/sparrovv/defme/configuration"
	"github.com/sparrovv/defme/hydra"
)

func main() {
	app := cli.NewApp()
	app.Name = "defme"
	app.Usage = "Get definion, transalation, synonyms and the usage of the word"
	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "s",
			Usage:     "Run HTTP server",
			Flags: []cli.Flag{
				cli.StringFlag{"port, p", "", "On what port it has to run"},
			},

			Action: func(c *cli.Context) {
				port := c.String("port")
				if len(port) == 0 {
					port = "8080"
				}
				config, err := configuration.FromEnv()
				if err != nil {
					log.Fatal(err)
					return
				}

				hydra.Serve(port, config)
			},
		},
		{
			Name:      "define",
			ShortName: "d",
			Usage:     "Type a word you want the definition and translation for",
			Flags: []cli.Flag{
				cli.StringFlag{"json, j", "", "Return response in json"},
				cli.StringFlag{"to, t", "", "Translate to your native language"},
			},

			Action: func(c *cli.Context) {
				config, err := configuration.FromEnv()
				if err != nil {
					log.Fatal(err)
					return
				}

				translateTo := c.String("to")

				toJSON := false
				if len(c.String("json")) != 0 {
					toJSON = true
				}

				term := strings.Join(c.Args(), " ")

				validateInput(term)
				fmt.Println(hydra.BuildResponse(term, config, translateTo, toJSON))
			},
		},
	}

	app.Run(os.Args)
}

func validateInput(input string) {
	if len(input) == 0 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
		return
	}
}
