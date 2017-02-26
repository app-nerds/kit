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
