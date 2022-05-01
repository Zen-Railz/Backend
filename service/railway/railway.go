package railway

import (
	"container/list"
	"fmt"
	"zenrailz/code"
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

func (s *Service) Journey(originStationName string, destinationStationName string) (interface{}, errorr.Entity) {
	//TODO: Remove identitycodemap - not-in-use
	repoStationNameMap, _, repoErr := s.railwayRepo.Network()
	if repoErr != nil {
		s.logger.Error(repoErr.Trace().Elaborate(), nil)
		return nil, repoErr
	}

	origin, originExist := repoStationNameMap[originStationName]
	if !originExist {
		return nil, errorr.New(code.RailwayServiceJourneyOriginNotFound, "Journey origin not found.", nil)
	}

	destination, destinationExist := repoStationNameMap[destinationStationName]
	if !destinationExist {
		return nil, errorr.New(code.RailwayServiceJourneyDestinationNotFound, "Journey destination not found.", nil)
	}

	journeys := [][]PathPoint{}
	itineraries := list.New()

	itinerary := Itinerary{
		Path: []PathPoint{
			{
				StationName:         origin.StationName,
				StationIdentityCode: "",
				StationLine:         "",
			},
		},
		Visited: map[string]struct{}{},
		Id:      "",
	}

	itineraries.PushBack(itinerary)

	s.logger.Info(fmt.Sprintf("PathQueue Length: %d", itineraries.Len()))

	for itineraries.Len() > 0 {
		itineraryItem := itineraries.Front()
		itineraries.Remove(itineraryItem)
		itinerary := itineraryItem.Value.(Itinerary)

		pathPoint := itinerary.Path[len(itinerary.Path)-1]

		_, hasVisited := itinerary.Visited[pathPoint.StationName]
		if hasVisited {
			continue
		} else {
			itinerary.Visited[pathPoint.StationName] = struct{}{}
		}

		if pathPoint.StationName == destination.StationName {
			s.logger.Info(fmt.Sprintf("(%s) Destination REACHED!", itinerary.Id))
			journeys = append(journeys, itinerary.Path)
		} else {
			s.logger.Info(fmt.Sprintf("(%s)[%s] %s - is not a destination", itinerary.Id, pathPoint.StationIdentityCode, pathPoint.StationName))
			stationNetworkNode := repoStationNameMap[pathPoint.StationName]

			// TODO: exclude those non-active stations

			for previousStationName := range stationNetworkNode.PreviousStationNames {
				newPath := make([]PathPoint, len(itinerary.Path))
				copy(newPath, itinerary.Path)
				newPath = append(newPath, PathPoint{
					StationName:         previousStationName,
					StationIdentityCode: "",
					StationLine:         "",
				})

				newVisited := make(map[string]struct{})
				for stationName, v := range itinerary.Visited {
					newVisited[stationName] = v
				}

				itineraries.PushBack(Itinerary{
					Path:    newPath,
					Visited: newVisited,
					Id:      itinerary.Id + "-" + previousStationName,
				})
			}

			for nextStationName := range stationNetworkNode.NextStationNames {
				newPath := make([]PathPoint, len(itinerary.Path))
				copy(newPath, itinerary.Path)
				newPath = append(newPath, PathPoint{
					StationName:         nextStationName,
					StationIdentityCode: "",
					StationLine:         "",
				})

				newVisited := make(map[string]struct{})
				for stationName, v := range itinerary.Visited {
					newVisited[stationName] = v
				}

				itineraries.PushBack(Itinerary{
					Path:    newPath,
					Visited: newVisited,
					Id:      itinerary.Id + "-" + nextStationName,
				})
			}
		}

		if len(journeys) == 3 {
			break
		}
	}

	s.logger.Info(fmt.Sprintf("Jouney Count: %d", len(journeys)))

	temp := struct {
		Journeys interface{}
		NameMap  interface{}
	}{
		Journeys: journeys,
		NameMap:  repoStationNameMap,
	}

	return temp, nil
}
