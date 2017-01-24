package output

import (
	"fmt"
	"os"

	"github.com/tscolari/plag/parser"
)

type Outputer interface {
	Write(data parser.Data) error
	Close() error
}

type Multi struct {
	outputers []Outputer
	open      bool
	channel   chan parser.Data
}

func NewMulti() *Multi {
	return &Multi{
		outputers: []Outputer{},
		open:      true,
	}
}

func (o *Multi) Add(outputer Outputer) {
	o.outputers = append(o.outputers, outputer)
}

func (o *Multi) Write(dataChan chan parser.Data) error {
	for data := range dataChan {
		for _, outputer := range o.outputers {
			if err := outputer.Write(data); err != nil {
				fmt.Fprintf(os.Stderr, "metric failed: %s", err.Error())
			}
		}
	}

	for _, outputer := range o.outputers {
		_ = outputer.Close()
	}

	return nil
}
