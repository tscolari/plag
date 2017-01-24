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
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return &Csv{file: file}, nil
}

func (o *Csv) Write(data parser.Data) error {
	fmt.Fprintf(o.file, "%s,%d\n", data.Timestamp.String(), data.Value)

	return nil
}

func (o *Csv) Close() error {
	return o.file.Close()
}
