package main

import (
	"os"
	"sync"

	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/parser"
	"github.com/urfave/cli"
)

type Outputer interface {
	Write(chan parser.Data) error
}

func main() {
	app := cli.NewApp()
	app.Name = "plag"
	app.Usage = "--message garden.streamIn [OPTIONS]"
	app.Description = "Plot lager time series"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "message",
			Usage: "message name to look for",
		},
		cli.StringFlag{
			Name:  "csv",
			Usage: "export csv file",
		},
		cli.StringFlag{
			Name:  "datadog-api-key",
			Usage: "datadog api key",
		},
		cli.StringFlag{
			Name:  "datadog-app-key",
			Usage: "datadog app key",
		},
	}

	app.Action = func(c *cli.Context) error {
		if !c.IsSet("message") {
			return cli.NewExitError("missing `--message` argument. Don't know what to look for.", 1)
		}

		parser := parser.New()
		data := parser.Parse(os.Stdin, c.String("message"))

		outputers := buildOutputers(c)
		var wg sync.WaitGroup
		for _, outputer := range outputers {
			wg.Add(1)
			go func(outputer Outputer) {
				_ = outputer.Write(data)
				wg.Done()
			}(outputer)
		}

		wg.Wait()
		return nil
	}

	_ = app.Run(os.Args)
}

func buildOutputers(c *cli.Context) []Outputer {
	outputers := []Outputer{}
	if c.IsSet("csv") {
		csv, _ := output.NewCsv(c.String("csv"))
		outputers = append(outputers, csv)
	}

	return outputers
}
