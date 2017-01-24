package parser_test

import (
	"bytes"

	"github.com/tscolari/plag/parser"
	"github.com/tscolari/plag/testassets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var subject *parser.Parser

	BeforeEach(func() {
		subject = parser.New()
	})

	Describe("Parse", func() {
		It("returns valid data points from the stream", func() {
			stream := bytes.NewBuffer([]byte(testassets.Log))

			dataChan := subject.Parse(stream, "test-app.function-1")
			expectedData := testassets.DataPointsPerMessage["test-app.function-1"]

			collectedData := []parser.Data{}
			for data := range dataChan {
				collectedData = append(collectedData, data)
			}
			Expect(collectedData).To(HaveLen(len(expectedData)))

			for _, data := range collectedData {
				Expect(expectedData).To(ContainElement(data))
			}
		})
	})

})
