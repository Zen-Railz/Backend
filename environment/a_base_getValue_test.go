package environment

import (
	"os"
	"zenrailz/anomaly"
	"zenrailz/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting environment variable", func() {
	Context("when value is set", func() {
		AfterEach(func() {
			os.Unsetenv(mockEnvironmentVariableKey)
		})

		It("should have value", func() {
			os.Setenv(mockEnvironmentVariableKey, mockEnvironmentVariableValue)
			value, _ := getValue(mockEnvironmentVariableKey)
			Expect(value).To(Equal(mockEnvironmentVariableValue))
		})
	})

	Context("when value is not set", func() {
		BeforeEach(func() {
			os.Unsetenv(mockEnvironmentVariableKey)
		})

		It("should not have value", func() {
			value, _ := getValue(mockEnvironmentVariableKey)
			Expect(value).To(BeEmpty(), "Value is %v.", value)
		})

		It("should have ServiceError", func() {
			_, err := getValue(mockEnvironmentVariableKey)
			serviceError, ok := err.(*anomaly.ServiceError)
			Expect(ok).To(BeTrue())
			Expect(serviceError.Code).To(Equal(code.EnvironmentVariableNotFound))
		})
	})
})
