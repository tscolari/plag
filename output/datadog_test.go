package output_test

import (
	"time"

	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/output/outputfakes"
	"github.com/tscolari/plag/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datadog", func() {
	var (
		datadog  *output.Datadog
		ddClient *outputfakes.FakeDatadogClient
	)

	BeforeEach(func() {
		ddClient = new(outputfakes.FakeDatadogClient)
		datadog = output.NewDatadog(ddClient, "my-metric")
	})

	It("sends the correct metric to datadog client", func() {
		data := parser.Data{
			Timestamp: time.Now(),
			Value:     5 * time.Minute,
		}

		Expect(datadog.Write(data)).To(Succeed())

		Expect(ddClient.PostMetricsCallCount()).To(Equal(1))

		metrics := ddClient.PostMetricsArgsForCall(0)
		Expect(metrics).To(HaveLen(1))

		Expect(metrics[0].Metric).To(Equal("my-metric"))
		Expect(metrics[0].Points).To(HaveLen(1))
		point := metrics[0].Points[0]

		Expect(point[0]).To(Equal(float64(data.Timestamp.Unix())))
		Expect(point[1]).To(Equal(float64(data.Value)))
	})
})
