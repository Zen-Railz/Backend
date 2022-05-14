package errorr

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Printing Error", func() {
	When("code and message exist", func() {
		It("should be well formatted", func() {
			err := &entity{
				code:    mockErrorCode,
				message: mockErrorMessage,
			}
			output := err.Error()
			Expect(output).To(Equal(fmt.Sprintf("(%s) %s", mockErrorCode, mockErrorMessage)))
		})
	})
})
