package logging_test

import (
	. "bitbucket.org/appninjas/kit/logging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LogFormat", func() {
	Context("the type LOG_FORMAT_SIMPLE,", func() {
		Context("when calling the String() method", func() {
			It("returns 'Simple'", func() {
				expected := "Simple"
				actual := LOG_FORMAT_SIMPLE.String()

				Expect(actual).To(Equal(expected))
			})
		})
	})

	Context("the type LOG_FORMAT_JSON,", func() {
		Context("when calling the String() method", func() {
			It("returns 'JSON'", func() {
				expected := "JSON"
				actual := LOG_FORMAT_JSON.String()

				Expect(actual).To(Equal(expected))
			})
		})
	})

	Context("The StringToLogFormat() method,", func() {
		Context("when given a valid format name", func() {
			It("returns a LogFormat type", func() {
				expected := "JSON"
				actual := StringToLogFormat("JSON").String()

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when given an invalid format name", func() {
			It("returns SimpleLogger", func() {
				expected := "Simple"
				actual := StringToLogFormat("TESTING").String()

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
