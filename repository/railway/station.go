package railway

import (
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

type Station struct {
	Name   string
	Code   string
	Number int
}

func (r *Repository) Stations() ([]Station, errorr.Entity) {
	stations := []Station{}

	rows, queryErr := r.database.Query("select name, code, number from station order by name, code, number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations.", queryErr)
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
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station.", scanErr)
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
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return stations, err.Trace()
	}

	return stations, nil
}
