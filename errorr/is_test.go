package errorr

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comparing error equality", func() {
	var errorrEntity error

	BeforeEach(func() {
		errorrEntity = &entity{
			code:    mockErrorCode,
			message: mockErrorMessage,
		}
	})

	Context("target error is an errorr entity", func() {
		When("code and message are same as errorr entity", func() {
			It("should be equal", func() {
				targetError := error(&entity{
					code:    mockErrorCode,
					message: mockErrorMessage,
				})
				output := errors.Is(errorrEntity, targetError)
				Expect(output).To(BeTrue())
			})
		})

		When("only the error code is the same as errorr entity", func() {
			It("should be equal", func() {
				targetError := error(&entity{
					code:    mockErrorCode,
					message: "different message",
				})
				output := errors.Is(errorrEntity, targetError)
				Expect(output).To(BeTrue())
			})
		})

		When("only message is the same as errorr entity", func() {
			It("should not be equal", func() {
				targetError := error(&entity{
					code:    "different code",
					message: mockErrorMessage,
				})
				output := errors.Is(errorrEntity, targetError)
				Expect(output).To(BeFalse())
			})
		})
	})

	Context("target is not an errorr entity", func() {
		It("should not be equal", func() {
			targetError := fmt.Errorf("(%s) %s", mockErrorCode, mockErrorMessage)
			output := errors.Is(errorrEntity, targetError)
			Expect(output).To(BeFalse())
		})
	})
})
