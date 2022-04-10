package railway

import (
	"zenrailz/anomaly"
)

func (s *Service) Stations() (interface{}, *anomaly.ServiceError) {
	dbStations, dbResErr := s.railwayRepo.Stations()
	if dbResErr != nil {
		dbResErr.Trace()
		s.logger.Error(dbResErr.Elaborate(), nil)
		return nil, dbResErr
	}

	stations := make(map[string]map[string]interface{})
	for _, dbStation := range dbStations {
		stationName := dbStation.Name
		stationCode := dbStation.Code
		stationNumber := dbStation.Number

		station, stationExist := stations[stationName]
		if stationExist {
			lineNumbers, lineExist := station[stationCode]
			if lineExist {
				switch v := lineNumbers.(type) {
				case int:
					station[stationCode] = []int{v, stationNumber}
				case []int:
					numberList := lineNumbers.([]int)
					numberList = append(numberList, stationNumber)
					station[stationCode] = numberList
				}
			} else {
				station[stationCode] = stationNumber
			}
		} else {
			newStation := make(map[string]interface{})
			newStation[stationCode] = stationNumber
			stations[stationName] = newStation
		}
	}

	return stations, nil
}
