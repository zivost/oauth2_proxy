package options

import (
	"bytes"
	"io"
	"testing"
	"time"

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

	BeforeEach(func() {
		// Make sure to clear previous test globals
		opts = nil
		err = nil
		config = nil
		configType = ""
		args = []string{}
	})

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

	Context("with a yaml configuration for cookies", func() {
		BeforeEach(func() {
			configType = "yaml"
			config = bytes.NewBuffer([]byte(`
        cookie:
          name: cookie_name
          secret: 123567890abcdef
          domain: example.com
          path: /path
          expire: 12h
          refresh: 1h
          secure: false
          httpOnly: false
        `))
		})

		It("sets the correct configuration", func() {
			expected := &CookieOptions{
				Name:     "cookie_name",
				Secret:   "123567890abcdef",
				Domain:   "example.com",
				Path:     "/path",
				Expire:   time.Duration(12) * time.Hour,
				Refresh:  time.Hour,
				Secure:   false,
				HTTPOnly: false,
			}
			Expect(opts.Cookie).To(Equal(expected))
		})

	})
})
