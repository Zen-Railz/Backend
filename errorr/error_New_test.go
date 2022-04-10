package errorr

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("New Error", func() {
	var err *entity

	When("error has only code and message", func() {
		BeforeEach(func() {
			err = New(mockErrorCode, mockErrorMessage, nil)
		})

		It("should have code", func() {
			Expect(err.code).To(Equal(mockErrorCode))
		})

		It("should have message", func() {
			Expect(err.message).To(Equal(mockErrorMessage))
		})

		It("should have history", func() {
			Expect(len(err.stackTrace)).To(Equal(1))
			Expect(err.stackTrace[0].FunctionName).NotTo(ContainSubstring("New"))
		})

		It("should not have annex", func() {
			Expect(err.annex).To(BeNil())
		})
	})
})
