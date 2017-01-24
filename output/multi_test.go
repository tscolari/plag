package output_test

import (
	"time"

	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/output/outputfakes"
	"github.com/tscolari/plag/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Multi", func() {
	Describe("Write", func() {
		var (
			multi     *output.Multi
			data      []parser.Data
			outputer1 *outputfakes.FakeOutputer
			outputer2 *outputfakes.FakeOutputer
			outputer3 *outputfakes.FakeOutputer
		)

		BeforeEach(func() {
			multi = output.NewMulti()
			data = []parser.Data{
				parser.Data{Timestamp: time.Now(), Value: 5 * time.Minute},
				parser.Data{Timestamp: time.Now(), Value: 10 * time.Minute},
				parser.Data{Timestamp: time.Now(), Value: 3 * time.Minute},
			}

			outputer1 = new(outputfakes.FakeOutputer)
			outputer2 = new(outputfakes.FakeOutputer)
			outputer3 = new(outputfakes.FakeOutputer)
		})

		It("sends the data to all outputers registed", func() {
			dataChan := make(chan parser.Data, 3)
			for _, d := range data {
				dataChan <- d
			}
			close(dataChan)

			multi.Add(outputer1)
			multi.Add(outputer2)
			multi.Add(outputer3)
			Expect(multi.Write(dataChan)).To(Succeed())

			Expect(outputer1.WriteCallCount()).To(Equal(len(data)))
			Expect(outputer2.WriteCallCount()).To(Equal(len(data)))
			Expect(outputer3.WriteCallCount()).To(Equal(len(data)))

			Expect(outputer1.WriteArgsForCall(0)).To(Equal(data[0]))
			Expect(outputer1.WriteArgsForCall(1)).To(Equal(data[1]))
			Expect(outputer1.WriteArgsForCall(2)).To(Equal(data[2]))

			Expect(outputer2.WriteArgsForCall(0)).To(Equal(data[0]))
			Expect(outputer2.WriteArgsForCall(1)).To(Equal(data[1]))
			Expect(outputer2.WriteArgsForCall(2)).To(Equal(data[2]))

			Expect(outputer3.WriteArgsForCall(0)).To(Equal(data[0]))
			Expect(outputer3.WriteArgsForCall(1)).To(Equal(data[1]))
			Expect(outputer3.WriteArgsForCall(2)).To(Equal(data[2]))

			Expect(outputer1.CloseCallCount()).To(Equal(1))
			Expect(outputer2.CloseCallCount()).To(Equal(1))
			Expect(outputer3.CloseCallCount()).To(Equal(1))
		})
	})
})
