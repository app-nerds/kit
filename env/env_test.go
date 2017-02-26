/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package env_test

import (
	"os"

	. "bitbucket.org/appninjas/kit/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env", func() {
	Describe("When getting an environment variable", func() {
		BeforeEach(func() {
			os.Unsetenv("KIT_ENV_TEST_VAR_1")
		})

		Context("that is not set", func() {
			It("should return the default value", func() {
				expected := "value 2"
				actual := Getenv("KIT_ENV_TEST_VAR_1", "value 2")

				Expect(actual).To(Equal(expected))
			})
		})

		Context("that is set", func() {
			BeforeEach(func() {
				os.Setenv("KIT_ENV_TEST_VAR_1", "value 1")
			})

			It("should return the value from the environment", func() {
				expected := "value 1"
				actual := Getenv("KIT_ENV_TEST_VAR_1", "value 1")

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
