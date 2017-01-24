package integration_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/tscolari/plag/testassets"
)

var _ = Describe("csv Flag", func() {
	var buffer *bytes.Buffer

	BeforeEach(func() {
		buffer = bytes.NewBuffer([]byte(testassets.Log))
	})

	It("writes the right data to the file", func() {
		tmpFile, err := ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())
		defer tmpFile.Close()

		cmd := exec.Command(PlagBin, "--message", "test-app.function-0", "--csv", tmpFile.Name())
		cmd.Stdin = buffer
		sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(sess).Should(gexec.Exit(0))

		data, err := ioutil.ReadAll(tmpFile)
		Expect(err).NotTo(HaveOccurred())

		dataPoints := testassets.DataPointsPerMessage["test-app.function-0"]
		for i := 0; i < len(dataPoints); i++ {
			Expect(string(data)).To(ContainSubstring(fmt.Sprintf("%s,%d\n", dataPoints[i].Timestamp.String(), dataPoints[i].Value)))
		}
	})
})
