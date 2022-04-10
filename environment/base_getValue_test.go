package environment

import (
	"os"
	"zenrailz/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting environment variable", func() {
	When("value is set", func() {
		AfterEach(func() {
			os.Unsetenv(mockEnvironmentVariableKey)
		})

		It("should have value", func() {
			os.Setenv(mockEnvironmentVariableKey, mockEnvironmentVariableValue)
			value, _ := getValue(mockEnvironmentVariableKey)
			Expect(value).To(Equal(mockEnvironmentVariableValue))
		})
	})

	When("value is not set", func() {
		BeforeEach(func() {
			os.Unsetenv(mockEnvironmentVariableKey)
			DeferCleanup(func() {
				os.Unsetenv(mockEnvironmentVariableKey)
			})
		})

		It("should not have value", func() {
			value, _ := getValue(mockEnvironmentVariableKey)
			Expect(value).To(BeEmpty(), "Value is %v.", value)
		})

		It("should have EnvironmentVariableNotFound status code", func() {
			_, err := getValue(mockEnvironmentVariableKey)
			Expect(err.Code()).To(Equal(code.EnvironmentVariableNotFound))
		})
	})
})
