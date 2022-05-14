package railway

import (
	"zenrailz/errorr"
)

func (s *Service) Stations() (interface{}, errorr.Entity) {
	repoStations, repoErr := s.railwayRepo.Stations()
	if repoErr != nil {
		repoErr.Trace()
		s.logger.Error(repoErr.Elaborate(), nil)
		return nil, repoErr
	}

	stations := make(map[string]map[string]interface{})
	for _, repoStation := range repoStations {
		stations[repoStation.Name] = make(map[string]interface{})

		for _, identity := range repoStation.Identifiers {
			numberContainer, prefixExist := stations[repoStation.Name][identity.Prefix]
			if prefixExist {
				switch v := numberContainer.(type) {
				case int:
					stations[repoStation.Name][identity.Prefix] = []int{v, identity.Number}
				case []int:
					numberList := numberContainer.([]int)
					stations[repoStation.Name][identity.Prefix] = append(numberList, identity.Number)
				}
			} else {
				stations[repoStation.Name][identity.Prefix] = identity.Number
			}
		}
	}

	return stations, nil
}

type StationIdentity struct {
	Code string
	Line string
}
