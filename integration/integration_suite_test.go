package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	PlagBin string
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		plagBin, err := gexec.Build("github.com/tscolari/plag")
		Expect(err).NotTo(HaveOccurred())

		return []byte(plagBin)
	}, func(data []byte) {
		PlagBin = string(data)
	})

	RunSpecs(t, "Integration Suite")
}
