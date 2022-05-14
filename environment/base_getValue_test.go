package environment

import (
	"os"
	"zenrailz/code"
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting environment variable", func() {
	When("value is set", func() {
		AfterEach(func() {
			os.Unsetenv(mock.EnvironmentVariableKey)
		})

		It("should have value", func() {
			os.Setenv(mock.EnvironmentVariableKey, mock.EnvironmentVariableValue)
			value, _ := getValue(mock.EnvironmentVariableKey)
			Expect(value).To(Equal(mock.EnvironmentVariableValue))
		})
	})

	When("value is not set", func() {
		BeforeEach(func() {
			os.Unsetenv(mock.EnvironmentVariableKey)
			DeferCleanup(func() {
				os.Unsetenv(mock.EnvironmentVariableKey)
			})
		})

		It("should not have value", func() {
			value, _ := getValue(mock.EnvironmentVariableKey)
			Expect(value).To(BeEmpty(), "Value is %v.", value)
		})

		It("should have EnvironmentVariableNotFound status code", func() {
			_, err := getValue(mock.EnvironmentVariableKey)
			Expect(err.Code()).To(Equal(code.EnvironmentVariableNotFound))
		})
	})
})
