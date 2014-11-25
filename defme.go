package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/sparrovv/defme/hydra"
	"github.com/sparrovv/defme/wordnik"
)

func main() {
	app := cli.NewApp()
	app.Name = "defme"
	app.Usage = "Get definion, transalation, synonyms and exampes of a word"
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
				client := wordnikClient()
				hydra.Serve(port, &client)
			},
		},
		{
			Name:      "define",
			ShortName: "d",
			Usage:     "define a word",
			Flags: []cli.Flag{
				cli.StringFlag{"json, j", "", "Return response in json"},
				cli.StringFlag{"to, t", "", "Translate to your native language"},
			},

			Action: func(c *cli.Context) {
				client := wordnikClient()

				translateTo := c.String("to")
				toJSON := false
				if len(c.String("json")) != 0 {
					toJSON = true
				}

				word := strings.Join(c.Args(), " ")
				validateInput(word)

				fmt.Println(hydra.BuildResponse(word, &client, translateTo, toJSON))
			},
		},
	}

	app.Run(os.Args)
}

func wordnikClient() wordnik.Client {
	wordnikApiKey := os.Getenv("WORDNIK_API_KEY")

	if len(wordnikApiKey) == 0 {
		panic("WORDNIK_API_KEY is not set")
	}

	return wordnik.NewClient(wordnikApiKey)
}

func validateInput(input string) {
	if len(input) == 0 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
		return
	}
}
