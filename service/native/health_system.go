package native

type SystemState struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

func (s *Service) SystemHealth() *SystemState {
	overallStatus := Healthy
	databaseHealth := s.DatabaseHealth()

	if databaseHealth.Status == Unhealthy {
		overallStatus = Unhealthy
	}

	return &SystemState{
		Status:   overallStatus,
		Database: databaseHealth.Status,
	}
}
