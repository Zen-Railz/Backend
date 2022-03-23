package health

const (
	Healthy   = "Healthy"
	Unhealthy = "Unhealthy"
)

type HolisticStatus struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

func Status() HolisticStatus {
	overallStatus := Healthy
	databaseStatus := Database()

	if databaseStatus.Status == Unhealthy {
		overallStatus = Unhealthy
	}

	return HolisticStatus{
		Status:   overallStatus,
		Database: databaseStatus.Status,
	}
}
