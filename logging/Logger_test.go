package logging_test

import (
	. "bitbucket.org/appninjas/kit/logging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	applicationName := "Kit Testing"

	Describe("LogFactory", func() {
		Context("when given a Simple Logger type", func() {
			It("should return a SimpleLogger object", func() {
				expected := true
				logger := LogFactory(LOG_FORMAT_SIMPLE, applicationName, FATAL)
				_, actual := logger.(*SimpleLogger)

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when given a JSON Logger type", func() {
			It("should return a JSONLogger object", func() {
				expected := true
				logger := LogFactory(LOG_FORMAT_JSON, applicationName, FATAL)
				_, actual := logger.(*JSONLogger)

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when given an invalid logger type", func() {
			It("should return a SimpleLogger object", func() {
				expected := true
				logger := LogFactory(-1, applicationName, FATAL)
				_, actual := logger.(*SimpleLogger)

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
