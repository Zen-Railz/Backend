package railway

import (
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting all stations from source", func() {
	var serviceUnderTest *Service

	When("failed to retrieve", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockConfigRepo := mock.NewConfigurationRepository()
			mockRailwayRepo := mock.NewRailwayRepository().
				SetSourceError()
			serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
		})

		It("should return error", func() {
			_, err := serviceUnderTest.Stations()
			Expect(err).To(HaveOccurred())
			Expect(err.Code()).To(Equal(mock.ErrorCode))
		})

		It("should not have any stations", func() {
			stations, _ := serviceUnderTest.Stations()
			Expect(stations).To(BeNil())
		})
	})
})

var _ = Describe("Structuring stations as output", func() {
	var serviceUnderTest *Service
	var mockLogger *mock.Logger
	var mockConfigRepo *mock.ConfigurationRepository
	var mockRailwayRepo *mock.RailwayRepository

	BeforeEach(func() {
		mockLogger = mock.NewLogger()
		mockConfigRepo = mock.NewConfigurationRepository()
	})

	When("source returns empty result", func() {
		BeforeEach(func() {
			mockRailwayRepo = mock.NewRailwayRepository().
				EmptyStations()
			serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
		})

		It("should have no stations", func() {
			stations, _ := serviceUnderTest.Stations()
			Expect(stations).To(BeEmpty())
		})

		It("should not have error", func() {
			_, err := serviceUnderTest.Stations()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("source returns stations", func() {
		var stations map[string]map[string]interface{}

		alphaStationName := "Station Alpha"
		bravoStationName := "Station Bravo"
		charlieStationName := "Station Charlie"
		deltaStationName := "Station Delta"
		echoStationName := "Station Echo"
		foxtrotStationName := "Station Foxtrot"

		northEastLineStationPrefix := "NE"
		eastWestLineStationPrefix := "EW"
		thomsonEastCoastLineStationPrefix := "TE"
		circleLineStationPrefix := "CC"

		BeforeEach(func() {
			mockRailwayRepo := mock.NewRailwayRepository().
				AddStation(alphaStationName, northEastLineStationPrefix, 1).
				AddStation(alphaStationName, northEastLineStationPrefix, 20).
				AddStation(bravoStationName, northEastLineStationPrefix, 2).
				AddStation(charlieStationName, northEastLineStationPrefix, 3).
				AddStation(deltaStationName, northEastLineStationPrefix, 4).
				AddStation(deltaStationName, eastWestLineStationPrefix, 15).
				AddStation(echoStationName, northEastLineStationPrefix, 5).
				AddStation(echoStationName, thomsonEastCoastLineStationPrefix, 7).
				AddStation(foxtrotStationName, northEastLineStationPrefix, 6).
				AddStation(foxtrotStationName, circleLineStationPrefix, 1).
				AddStation(foxtrotStationName, circleLineStationPrefix, 30)
			serviceUnderTest = NewService(mockLogger, mockConfigRepo, mockRailwayRepo)
			rawStations, _ := serviceUnderTest.Stations()
			stations, _ = rawStations.(map[string]map[string]interface{})
		})

		It("should have stations total count same as total number of unique station names", func() {
			Expect(len(stations)).To(Equal(6))
		})

		It("should output stations with multiple stops on the same line as an array", func() {
			alphaStationNumbers := stations[alphaStationName][northEastLineStationPrefix].([]int)
			Expect(len(alphaStationNumbers)).To(Equal(2))

			foxtrotStationNumbers := stations[foxtrotStationName][circleLineStationPrefix].([]int)
			Expect(len(foxtrotStationNumbers)).To(Equal(2))
		})

		It("should output stations with single stop as the station number", func() {
			bravoStationNumber := stations[bravoStationName][northEastLineStationPrefix].(int)
			Expect(bravoStationNumber).To(Equal(2))

			charlieStationNumber := stations[charlieStationName][northEastLineStationPrefix].(int)
			Expect(charlieStationNumber).To(Equal(3))

			deltaNorthEastStationNumber := stations[deltaStationName][northEastLineStationPrefix].(int)
			Expect(deltaNorthEastStationNumber).To(Equal(4))

			deltaEastWestStationNumber := stations[deltaStationName][eastWestLineStationPrefix].(int)
			Expect(deltaEastWestStationNumber).To(Equal(15))

			echoNorthEastStationNumber := stations[echoStationName][northEastLineStationPrefix].(int)
			Expect(echoNorthEastStationNumber).To(Equal(5))

			echoThomsonEastCoastStationNumber := stations[echoStationName][thomsonEastCoastLineStationPrefix].(int)
			Expect(echoThomsonEastCoastStationNumber).To(Equal(7))

			foxtrotNorthEastStationNumber := stations[foxtrotStationName][northEastLineStationPrefix].(int)
			Expect(foxtrotNorthEastStationNumber).To(Equal(6))
		})

	})

})
