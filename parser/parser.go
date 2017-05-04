package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"code.cloudfoundry.org/lager/chug"
)

var validStarters = []string{"begin", "start", "starting", "started"}
var validFinishers = []string{"finish", "end", "ending", "finished", "done"}

type Data struct {
	Timestamp time.Time
	Value     time.Duration
}

type Parser struct {
	dataChan    chan Data
	pendingData map[string]Data
}

func New() *Parser {
	return &Parser{
		dataChan:    make(chan Data),
		pendingData: map[string]Data{},
	}
}

func (p *Parser) Parse(input io.Reader, message string) chan Data {
	scanner := bufio.NewReader(input)
	go func() {
		for {
			line, err := scanner.ReadBytes('\n')
			if line != nil {
				if strings.Contains(string(line), message) {
					p.verifyEntry(message, line)
				}
			}

			if err != nil {
				close(p.dataChan)
				break
			}
		}
	}()

	return p.dataChan
}

func (p *Parser) verifyEntry(message string, rawEntry []byte) {
	entry := entry(rawEntry)
	if !entry.IsLager {
		return
	}

	if p.isStarter(message, entry) {
		p.pendingData[entry.Log.Session] = Data{Timestamp: entry.Log.Timestamp}
		return
	}

	if p.isFinisher(message, entry) {
		data, ok := p.pendingData[entry.Log.Session]
		if !ok {
			return
		}

		data.Value = entry.Log.Timestamp.Sub(data.Timestamp)
		p.dataChan <- data

		delete(p.pendingData, entry.Log.Session)
		return
	}
}

func (p *Parser) isStarter(message string, entry chug.Entry) bool {
	for _, starter := range validStarters {
		if fmt.Sprintf("%s.%s", message, starter) == entry.Log.Message {
			return true
		}
	}

	return false
}

func (p *Parser) isFinisher(message string, entry chug.Entry) bool {
	for _, finisher := range validFinishers {
		if fmt.Sprintf("%s.%s", message, finisher) == entry.Log.Message {
			return true
		}
	}

	return false
}
