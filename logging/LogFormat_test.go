/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package logging_test

import (
	. "code.appninjas.biz/appninjas/kit/logging"

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
