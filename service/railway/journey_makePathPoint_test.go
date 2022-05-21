package railway

import (
	"zenrailz/code"
	"zenrailz/mock"
	"zenrailz/repository/railway"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a node in the railway network", func() {
	var serviceUnderTest *Service
	var networkNode *railway.NetworkNode
	var mockLogger *mock.Logger
	var mockConfigRepo *mock.ConfigurationRepository
	var mockRailwayRepo *mock.RailwayRepository

	BeforeEach(func() {
		mockLogger = mock.NewLogger()
		mockConfigRepo = mock.NewConfigurationRepository()
		mockRailwayRepo = mock.NewRailwayRepository()
		serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
		networkNode = &railway.NetworkNode{
			StationName: mock.RailwayStationName,
		}
	})

	When("at least one identity is active", func() {
		numActiveIdentities := 2

		BeforeEach(func() {
			networkNode.StationIdentities = make(map[string]railway.StationIdentity)

			networkNode.StationIdentities["activeIdentityOne"] = railway.StationIdentity{
				IsActive: true,
			}
			networkNode.StationIdentities["activeIdentityTwo"] = railway.StationIdentity{
				IsActive: true,
			}
			networkNode.StationIdentities["inactiveIdentity"] = railway.StationIdentity{
				IsActive: false,
			}
		})

		It("should not have error", func() {
			_, err := serviceUnderTest.makePathPoint(networkNode)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should have a valid name", func() {
			point, _ := serviceUnderTest.makePathPoint(networkNode)
			Expect(point.StationName).To(Equal(mock.RailwayStationName))
		})

		It("should contain all active identites only", func() {
			point, _ := serviceUnderTest.makePathPoint(networkNode)
			Expect(len(point.StationIdentities)).To(Equal(numActiveIdentities))
		})
	})

	When("there are no active identities", func() {
		BeforeEach(func() {
			networkNode.StationIdentities = make(map[string]railway.StationIdentity)
			networkNode.StationIdentities["inactiveIdentity"] = railway.StationIdentity{
				IsActive: false,
			}
		})

		It("should have error", func() {
			_, err := serviceUnderTest.makePathPoint(networkNode)
			Expect(err).To(HaveOccurred())
			Expect(err.Code()).To(Equal(code.RailwayServiceStationUnavailable))
		})
	})
})
