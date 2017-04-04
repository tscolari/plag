package output

import (
	"github.com/tscolari/plag/parser"
	datadog "github.com/zorkian/go-datadog-api"
)

type DatadogClient interface {
	PostMetrics(series []datadog.Metric) error
}

type Datadog struct {
	client     DatadogClient
	metricName string
}

func NewDatadog(client DatadogClient, metricName string) *Datadog {
	return &Datadog{
		client:     client,
		metricName: metricName,
	}
}

func (o *Datadog) Write(data parser.Data) error {
	metric := datadog.Metric{
		Metric: &o.metricName,
		Points: []datadog.DataPoint{
			datadog.DataPoint{
				float64(data.Timestamp.Unix()),
				float64(data.Value),
			},
		},
	}

	return o.client.PostMetrics([]datadog.Metric{metric})
}

func (o *Datadog) Close() error {
	return nil
}
