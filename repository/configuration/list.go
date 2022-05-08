package configuration

import (
	"fmt"
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func (r *Repository) List(category string) (map[string]string, errorr.Entity) {
	configuration := make(map[string]string)

	rows, queryErr := r.database.Query("select c.key, c.value from configuration c where category=$1 and is_deleted=false", category)
	if queryErr != nil {
		message := fmt.Sprintf("Unable to get configuration for %s from database.", category)
		err := common.ParseError(code.DatabaseQueryFailure, message, queryErr)
		return configuration, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value string

		if scanErr := rows.Scan(&key, &value); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a key from database.", scanErr)
			return configuration, err.Trace()
		}

		configuration[key] = value
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return configuration, err.Trace()
	}

	return configuration, nil
}
