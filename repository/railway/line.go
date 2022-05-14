package railway

import (
	"database/sql"
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func (r *Repository) Lines() ([]Line, errorr.Entity) {
	lines := []Line{}

	rows, queryErr := r.database.Query("select name, code, type, is_active, announcement from line order by name")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get lines.", queryErr)
		return lines, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		var lineName string
		var lineCode string
		var lineType string
		var lineIsActive bool
		var lineAnnouncement sql.NullString

		if scanErr := rows.Scan(&lineName, &lineCode, &lineType, &lineIsActive, &lineAnnouncement); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a line.", scanErr)
			return lines, err.Trace()
		}

		lines = append(lines, Line{
			Name:         lineName,
			Code:         lineCode,
			Type:         lineType,
			IsActive:     lineIsActive,
			Announcement: lineAnnouncement.String,
		})
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return lines, err.Trace()
	}

	return lines, nil
}

type Line struct {
	Name         string
	Code         string
	Type         string
	IsActive     bool
	Announcement string
}
