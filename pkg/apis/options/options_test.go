package options

import (
	"io"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOptions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Options")
}

var _ = Describe("Load", func() {
	var opts *Options
	var err error
	var config io.Reader
	var configType string
	var args []string

	JustBeforeEach(func() {
		opts, err = Load(config, configType, args)
	})

	Context("with no configuration", func() {
		It("returns no error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the default configuration", func() {
			defaultOpts := New()
			Expect(opts).To(Equal(defaultOpts))
		})
	})
})
