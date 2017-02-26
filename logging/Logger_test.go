/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
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
