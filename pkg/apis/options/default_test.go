package options

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("defaultStruct", func() {
	type InnerTest struct {
		TestString string `default:"bar"`
	}

	type Test struct {
		TestString    string        `default:"foo"`
		TestBool      bool          `default:"true"`
		TestDuration  time.Duration `default:"60000000000"`
		TestPtrStruct *InnerTest
		TestStruct    InnerTest
	}

	var testStruct *Test

	BeforeEach(func() {
		testStruct = &Test{}
		err := defaultStruct(testStruct)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("with a string field", func() {
		It("reads the correct default", func() {
			Expect(testStruct.TestString).To(Equal("foo"))
		})
	})

	Context("with a bool field", func() {
		It("reads the correct default", func() {
			Expect(testStruct.TestBool).To(Equal(true))
		})
	})

	Context("with a duration field", func() {
		It("reads the correct default", func() {
			Expect(testStruct.TestDuration).To(Equal(time.Minute))
		})
	})

	Context("with a pointer struct field", func() {
		It("defaults the values in the struct", func() {
			Expect(testStruct.TestPtrStruct.TestString).To(Equal("bar"))
		})
	})

	Context("with a struct field", func() {
		It("defaults the values in the struct", func() {
			Expect(testStruct.TestStruct.TestString).To(Equal("bar"))
		})
	})
})
