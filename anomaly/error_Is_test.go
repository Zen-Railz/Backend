package anomaly

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comparing ServiceError equality", func() {
	var serviceError error

	BeforeEach(func() {
		serviceError = &ServiceError{
			Code:    mockServiceErrorCode,
			Message: mockServiceErrorMessage,
		}
	})

	Context("target error is a ServiceError", func() {
		When("code and message are same as service error", func() {
			It("should be equal", func() {
				err := error(&ServiceError{
					Code:    mockServiceErrorCode,
					Message: mockServiceErrorMessage,
				})
				output := errors.Is(serviceError, err)
				Expect(output).To(BeTrue())
			})
		})

		When("only code is the same as service error", func() {
			It("should be equal", func() {
				err := error(&ServiceError{
					Code:    mockServiceErrorCode,
					Message: "different message",
				})
				output := errors.Is(serviceError, err)
				Expect(output).To(BeTrue())
			})
		})

		When("only message is the same as service error", func() {
			It("should not be equal", func() {
				err := error(&ServiceError{
					Code:    "different code",
					Message: mockServiceErrorMessage,
				})
				output := errors.Is(serviceError, err)
				Expect(output).To(BeFalse())
			})
		})
	})

	Context("target is not a ServiceError", func() {
		It("should not be equal", func() {
			err := fmt.Errorf("(%s) %s", mockServiceErrorCode, mockServiceErrorMessage)
			output := errors.Is(serviceError, err)
			Expect(output).To(BeFalse())
		})
	})
})
