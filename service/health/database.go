package health

import (
	"zenrailz/data/database"
	"zenrailz/log"
)

type DatabaseStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func Database() DatabaseStatusResponse {
	db, conErr := database.Connect()
	if conErr != nil {
		conErr.Trace()
		log.Error(conErr.Elaborate(), nil)
		return DatabaseStatusResponse{
			Status:  Unhealthy,
			Message: conErr.Elaborate(),
		}
	}
	defer db.Close()

	healthErr := database.Health(db)
	if healthErr != nil {
		healthErr.Trace()
		log.Error(healthErr.Elaborate(), nil)
		return DatabaseStatusResponse{
			Status:  Unhealthy,
			Message: healthErr.Elaborate(),
		}
	}

	return DatabaseStatusResponse{
		Status: Healthy,
	}
}
