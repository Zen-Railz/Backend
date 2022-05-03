package railway

import (
	"container/list"
	"fmt"
	"strconv"
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/railway"
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

func (s *Service) Journey(originStationName string, destinationStationName string) ([][]PathPoint, errorr.Entity) {
	repoStationNameMap, repoErr := s.railwayRepo.Network()
	if repoErr != nil {
		s.logger.Error(repoErr.Trace().Elaborate(), nil)
		return nil, repoErr
	}

	destination, destinationExist := repoStationNameMap[destinationStationName]
	if !destinationExist {
		return nil, errorr.New(code.RailwayServiceJourneyDestinationNotFound, "Journey destination not found.", nil)
	}

	origin, originExist := repoStationNameMap[originStationName]
	if !originExist {
		return nil, errorr.New(code.RailwayServiceJourneyOriginNotFound, "Journey origin not found.", nil)
	}

	originPathPoint, err := s.makePathPoint(origin)
	if err != nil {
		s.logger.Error(err.Trace().Elaborate(), nil)
		return nil, err
	}

	journey := [][]PathPoint{}
	bfeQueue := list.New()

	bfeQueue.PushBack(BfeQueueObject{
		Path:    []PathPoint{originPathPoint},
		Visited: map[string]struct{}{},
	})

	for bfeQueue.Len() > 0 {
		bfeItemPtr := bfeQueue.Front()
		bfeQueue.Remove(bfeItemPtr)

		bfeItem := bfeItemPtr.Value.(BfeQueueObject)
		pathPoint := bfeItem.Path[len(bfeItem.Path)-1]

		_, hasVisited := bfeItem.Visited[pathPoint.StationName]
		if hasVisited {
			continue
		} else {
			// Mark station as visited
			bfeItem.Visited[pathPoint.StationName] = struct{}{}
		}

		if pathPoint.StationName == destination.StationName {
			journey = append(journey, bfeItem.Path)
		} else {
			stationNetworkNode := repoStationNameMap[pathPoint.StationName]

			for _, nextNode := range stationNetworkNode.AdjacentNodes {
				pathPoint, err := s.makePathPoint(nextNode)
				if err != nil {
					s.logger.Error(err.Trace().Elaborate(), nil)
					return nil, err
				}

				newPath := make([]PathPoint, len(bfeItem.Path))
				copy(newPath, bfeItem.Path)
				newPath = append(newPath, pathPoint)

				newVisited := make(map[string]struct{})
				for stationName := range bfeItem.Visited {
					newVisited[stationName] = struct{}{}
				}

				bfeQueue.PushBack(BfeQueueObject{
					Path:    newPath,
					Visited: newVisited,
				})
			}
		}

		if len(journey) == 3 {
			break
		}
	}

	return journey, nil
}

func (s *Service) makePathPoint(networkNode *railway.NetworkNode) (PathPoint, errorr.Entity) {
	point := PathPoint{
		StationName: networkNode.StationName,
	}

	for _, identity := range networkNode.StationIdentities {
		if identity.IsActive {
			point.StationIdentities = append(point.StationIdentities, StationIdentity{
				Code: identity.Prefix + strconv.Itoa(identity.Number),
				Line: identity.Line,
			})
		}
	}

	if len(point.StationIdentities) == 0 {
		message := fmt.Sprintf("[%s] station is currently undergoing maintenance.", point.StationName)
		return point, errorr.New(code.RailwayServiceStationUnavailable, message, nil)
	}

	return point, nil
}
