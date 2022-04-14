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
			mockRailwayRepo := mock.NewRailwayRepository().
				SetSourceError()
			serviceUnderTest = NewService(mockLogger, mockRailwayRepo)
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
	var mockRailwayRepo *mock.RailwayRepository

	When("source returns empty result", func() {
		BeforeEach(func() {
			mockLogger = mock.NewLogger()
			mockRailwayRepo = mock.NewRailwayRepository().
				EmptyStations()
			serviceUnderTest = NewService(mockLogger, mockRailwayRepo)
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

		mockLogger := mock.NewLogger()

		alphaStationName := "Station Alpha"
		bravoStationName := "Station Bravo"
		charlieStationName := "Station Charlie"
		deltaStationName := "Station Delta"
		echoStationName := "Station Echo"
		foxtrotStationName := "Station Foxtrot"

		northEastLineCode := "NEL"
		eastWestLineCode := "EWL"
		thomsonEastCoastLineCode := "TEL"
		circleLineCode := "CCL"

		BeforeEach(func() {
			mockRailwayRepo := mock.NewRailwayRepository().
				EmptyStations().
				AddStation(alphaStationName, northEastLineCode, 1).
				AddStation(alphaStationName, northEastLineCode, 20).
				AddStation(bravoStationName, northEastLineCode, 2).
				AddStation(charlieStationName, northEastLineCode, 3).
				AddStation(deltaStationName, northEastLineCode, 4).
				AddStation(deltaStationName, eastWestLineCode, 15).
				AddStation(echoStationName, northEastLineCode, 5).
				AddStation(echoStationName, thomsonEastCoastLineCode, 7).
				AddStation(foxtrotStationName, northEastLineCode, 6).
				AddStation(foxtrotStationName, circleLineCode, 1).
				AddStation(foxtrotStationName, circleLineCode, 30)
			serviceUnderTest = NewService(mockLogger, mockRailwayRepo)
			rawStations, _ := serviceUnderTest.Stations()
			stations, _ = rawStations.(map[string]map[string]interface{})
		})

		It("should have stations total count same as total number of unqie station names", func() {
			Expect(len(stations)).To(Equal(6))
		})

		It("should output stations with multiple stops on the same line as an array", func() {
			alphaStationNumbers := stations[alphaStationName][northEastLineCode].([]int)
			Expect(len(alphaStationNumbers)).To(Equal(2))

			foxtrotStationNumbers := stations[foxtrotStationName][circleLineCode].([]int)
			Expect(len(foxtrotStationNumbers)).To(Equal(2))
		})

		It("should output stations with single stop as the station number", func() {
			bravoStationNumber := stations[bravoStationName][northEastLineCode].(int)
			Expect(bravoStationNumber).To(Equal(2))

			charlieStationNumber := stations[charlieStationName][northEastLineCode].(int)
			Expect(charlieStationNumber).To(Equal(3))

			deltaNorthEastStationNumber := stations[deltaStationName][northEastLineCode].(int)
			Expect(deltaNorthEastStationNumber).To(Equal(4))

			deltaEastWestStationNumber := stations[deltaStationName][eastWestLineCode].(int)
			Expect(deltaEastWestStationNumber).To(Equal(15))

			echoNorthEastStationNumber := stations[echoStationName][northEastLineCode].(int)
			Expect(echoNorthEastStationNumber).To(Equal(5))

			echoThomsonEastCoastStationNumber := stations[echoStationName][thomsonEastCoastLineCode].(int)
			Expect(echoThomsonEastCoastStationNumber).To(Equal(7))

			foxtrotNorthEastStationNumber := stations[foxtrotStationName][northEastLineCode].(int)
			Expect(foxtrotNorthEastStationNumber).To(Equal(6))
		})

	})

})
