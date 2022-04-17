package railway

import "zenrailz/errorr"

func (s *Service) Stations() (interface{}, errorr.Entity) {
	dbStations, dbResErr := s.railwayRepo.Stations()
	if dbResErr != nil {
		dbResErr.Trace()
		s.logger.Error(dbResErr.Elaborate(), nil)
		return nil, dbResErr
	}

	stations := make(map[string]map[string]interface{})
	for _, dbStation := range dbStations {
		stationName := dbStation.Name
		stationPrefix := dbStation.Prefix
		stationNumber := dbStation.Number

		station, stationExist := stations[stationName]
		if stationExist {
			lineNumbers, lineExist := station[stationPrefix]
			if lineExist {
				switch v := lineNumbers.(type) {
				case int:
					station[stationPrefix] = []int{v, stationNumber}
				case []int:
					numberList := lineNumbers.([]int)
					numberList = append(numberList, stationNumber)
					station[stationPrefix] = numberList
				}
			} else {
				station[stationPrefix] = stationNumber
			}
		} else {
			newStation := make(map[string]interface{})
			newStation[stationPrefix] = stationNumber
			stations[stationName] = newStation
		}
	}

	return stations, nil
}
