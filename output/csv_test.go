package output_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csv", func() {
	var (
		outputer   *output.Csv
		outputFile *os.File
	)

	BeforeEach(func() {
		var err error
		outputFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())
		outputer, err = output.NewCsv(outputFile.Name())
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Write", func() {
		var data []parser.Data

		BeforeEach(func() {
			data = []parser.Data{
				parser.Data{Timestamp: time.Now(), Value: 5 * time.Minute},
				parser.Data{Timestamp: time.Now(), Value: 10 * time.Minute},
				parser.Data{Timestamp: time.Now(), Value: 3 * time.Minute},
			}
		})

		It("writes the data in a csv format", func() {
			for _, d := range data {
				Expect(outputer.Write(d)).To(Succeed())
			}
			writenData, err := ioutil.ReadAll(outputFile)
			Expect(err).NotTo(HaveOccurred())
			for _, d := range data {
				Expect(string(writenData)).To(ContainSubstring(fmt.Sprintf("%s,%d\n", d.Timestamp.String(), d.Value)))
			}
		})
	})
})
