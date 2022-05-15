package railway

import (
	"zenrailz/code"
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting a list of journey between two locations", func() {
	var serviceUnderTest *Service
	var mockLogger *mock.Logger
	var mockConfigRepo *mock.ConfigurationRepository
	var mockRailwayRepo *mock.RailwayRepository

	BeforeEach(func() {
		mockLogger = mock.NewLogger()
		mockConfigRepo = mock.NewConfigurationRepository()
	})

	Context("Retrieving station map from source", func() {
		When("failed to retrieve", func() {
			BeforeEach(func() {
				mockRailwayRepo = mock.NewRailwayRepository().
					SetSourceError()
				serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
			})

			It("should return error", func() {
				_, err := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
				Expect(err).To(HaveOccurred())
				Expect(err.Code()).To(Equal(mock.ErrorCode))
			})

			It("should not have any journey", func() {
				journey, _ := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
				Expect(journey).To(BeNil())
			})
		})
	})

	When("invalid destination is provided", func() {
		BeforeEach(func() {
			mockRailwayRepo = mock.NewRailwayRepository()
			serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
		})

		It("should return error", func() {
			_, err := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
			Expect(err).To(HaveOccurred())
			Expect(err.Code()).To(Equal(code.RailwayServiceJourneyDestinationNotFound))
		})

		It("should not have any result", func() {
			journey, _ := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
			Expect(journey).To(BeNil())
		})
	})

	When("invalid origin is provided", func() {
		BeforeEach(func() {
			mockRailwayRepo = mock.NewRailwayRepository().
				AddNetworkNode(mock.RailwayJourneyDestination, nil, nil)
			serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
		})

		It("should return error", func() {
			_, err := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
			Expect(err).To(HaveOccurred())
			Expect(err.Code()).To(Equal(code.RailwayServiceJourneyOriginNotFound))
		})

		It("should not have any result", func() {
			journey, _ := serviceUnderTest.Journey(mock.RailwayJourneyOrigin, mock.RailwayJourneyDestination)
			Expect(journey).To(BeNil())
		})
	})
})
