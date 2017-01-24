package output

import (
	"fmt"
	"os"

	"github.com/tscolari/plag/parser"
)

type Csv struct {
	file *os.File
}

func NewCsv(filePath string) (*Csv, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return &Csv{file: file}, nil
}

func (o *Csv) Write(dataChan chan parser.Data) error {
	for data := range dataChan {
		fmt.Fprintf(o.file, "%s,%d\n", data.Timestamp.String(), data.Value)
	}

	return o.file.Close()
}
