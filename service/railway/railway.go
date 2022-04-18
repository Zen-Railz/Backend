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

func (s *Service) Lines() ([]Line, errorr.Entity) {
	dbLines, dbResErr := s.railwayRepo.Lines()
	if dbResErr != nil {
		dbResErr.Trace()
		s.logger.Error(dbResErr.Elaborate(), nil)
		return nil, dbResErr
	}

	lines := []Line{}
	for _, dbLine := range dbLines {
		lines = append(lines, Line{
			Name:         dbLine.Name,
			Code:         dbLine.Code,
			Type:         dbLine.Type,
			IsActive:     dbLine.IsActive,
			Announcement: dbLine.Announcement,
		})
	}

	return lines, nil
}
