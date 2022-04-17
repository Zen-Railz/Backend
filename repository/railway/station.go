package railway

import (
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func (r *Repository) Stations() ([]Station, errorr.Entity) {
	stations := []Station{}

	rows, queryErr := r.database.Query("select s.name, s.prefix, s.number, s.is_active, l.name from station s, line l where s.line = l.code order by s.name, s.prefix, s.number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations.", queryErr)
		return stations, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		station := Station{}

		if scanErr := rows.Scan(&station.Name, &station.Prefix, &station.Number, &station.IsActive, &station.Line); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station.", scanErr)
			return stations, err.Trace()
		}

		stations = append(stations, station)
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return stations, err.Trace()
	}

	return stations, nil
}
