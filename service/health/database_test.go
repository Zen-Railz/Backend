package health

import (
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checking the health of the database", func() {
	var serviceUnderTest *Service

	When("able to ping database", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockDbRepo := mock.NewDatabaseRepository()
			serviceUnderTest = NewService(mockLogger, mockDbRepo)
		})

		It("should return status as healthy", func() {
			health := serviceUnderTest.Database()
			Expect(health.Status).To(Equal(Healthy))
		})

		It("should not return any message", func() {
			health := serviceUnderTest.Database()
			Expect(health.Message).To(BeEmpty())
		})
	})

	When("failed to ping database", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockDbRepo := mock.NewDatabaseRepository().
				SetPingError()
			serviceUnderTest = NewService(mockLogger, mockDbRepo)
		})

		It("should return status as unhealthy", func() {
			health := serviceUnderTest.Database()
			Expect(health.Status).To(Equal(Unhealthy))
		})

		It("should return error message", func() {
			health := serviceUnderTest.Database()
			Expect(health.Message).To(Equal(mock.ErrorStackTraceMessage))
		})
	})
})
