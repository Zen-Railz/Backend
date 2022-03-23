package railway

import (
	"zenrailz/anomaly"
	"zenrailz/data/database"
	"zenrailz/log"
)

func Stations() (interface{}, *anomaly.ServiceError) {

	db, dbConErr := database.Connect()
	if dbConErr != nil {
		dbConErr.Trace()
		log.Error(dbConErr.Elaborate(), nil)
		return nil, dbConErr
	}
	defer db.Close()

	dbStations, dbResErr := database.GetStations(db)
	if dbResErr != nil {
		dbResErr.Trace()
		log.Error(dbResErr.Elaborate(), nil)
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
