package anomaly

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Printing Error", func() {
	When("code and message exist", func() {
		It("should be well formatted", func() {
			err := &ServiceError{
				Code:    mockServiceErrorCode,
				Message: mockServiceErrorMessage,
			}
			output := err.Error()
			Expect(output).To(Equal(fmt.Sprintf("(%s) %s", mockServiceErrorCode, mockServiceErrorMessage)))
		})
	})
})
