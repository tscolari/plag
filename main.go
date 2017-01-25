package main

import (
	"os"

	"github.com/gizak/termui"
	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/parser"
	"github.com/urfave/cli"
	datadog "github.com/zorkian/go-datadog-api"
)

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
		cli.StringFlag{
			Name:  "datadog-metric-name",
			Usage: "metric name to post to datadog",
		},
		cli.BoolFlag{
			Name:  "graph",
			Usage: "plot realtime graph on terminal",
		},
	}

	app.Action = func(c *cli.Context) error {
		if !c.IsSet("message") {
			return cli.NewExitError("missing `--message` argument. Don't know what to look for.", 1)
		}

		parser := parser.New()
		data := parser.Parse(os.Stdin, c.String("message"))

		if c.IsSet("graph") {
			if err := termui.Init(); err != nil {
				panic(err)
			}
			defer termui.Close()

			termui.Handle("/sys/kbd/q", func(termui.Event) {
				close(data)
				termui.StopLoop()
			})
		}

		outputer := output.NewMulti()
		addOutputers(outputer, c)
		_ = outputer.Write(data)

		if c.IsSet("graph") {
			termui.Loop()
		}
		return nil
	}

	_ = app.Run(os.Args)
}

func addOutputers(multi *output.Multi, c *cli.Context) {
	if c.IsSet("csv") {
		csv, _ := output.NewCsv(c.String("csv"))
		multi.Add(csv)
	}

	if c.IsSet("datadog-api-key") {
		if !c.IsSet("datadog-app-key") {
			panic("Missing `--datadog-app-key` flag")
		}

		client := datadog.NewClient(c.String("datadog-api-key"), c.String("datadog-app-key"))
		metricName := c.String("message")
		if c.IsSet("datadog-metric-name") {
			metricName = c.String("datadog-metric-name")
		}
		datadog := output.NewDatadog(client, metricName)
		multi.Add(datadog)
	}

	if c.IsSet("graph") {
		graph := output.NewGraph(c.String("message"))
		multi.Add(graph)
	}
}
