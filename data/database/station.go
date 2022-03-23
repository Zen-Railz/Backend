package database

import (
	"database/sql"
	"zenrailz/anomaly"
	"zenrailz/code"
)

type Station struct {
	Name   string
	Code   string
	Number int
}

func GetStations(db *sql.DB) ([]Station, *anomaly.ServiceError) {
	stations := []Station{}

	rows, queryErr := db.Query("select name, code, number from station order by name, code, number")
	if queryErr != nil {
		err := parseError(code.DatabaseQueryFailure, "Unable to get stations.", queryErr)
		return stations, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		var (
			stationName   string
			stationCode   string
			stationNumber int
		)

		if scanErr := rows.Scan(&stationName, &stationCode, &stationNumber); scanErr != nil {
			err := parseError(code.DatabaseRowScanFailure, "Unable to read a station.", scanErr)
			return stations, err.Trace()
		}

		station := &Station{
			Name:   stationName,
			Code:   stationCode,
			Number: stationNumber,
		}

		stations = append(stations, *station)
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := parseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return stations, err.Trace()
	}

	return stations, nil
}
