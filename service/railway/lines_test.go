package railway

import (
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting all lines from source", func() {
	var serviceUnderTest *Service

	When("failed to retrieve", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockRailwayRepo := mock.NewRailwayRepository().
				SetSourceError()
			serviceUnderTest = NewService(mockLogger, mockRailwayRepo)
		})

		It("should return error", func() {
			_, err := serviceUnderTest.Lines()
			Expect(err).To(HaveOccurred())
			Expect(err.Code()).To(Equal(mock.ErrorCode))
		})

		It("should not have any lines", func() {
			lines, _ := serviceUnderTest.Lines()
			Expect(lines).To(BeNil())
		})
	})
})
