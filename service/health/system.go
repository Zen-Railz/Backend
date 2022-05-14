package health

func (s *Service) System() *SystemState {
	overallStatus := Healthy
	databaseHealth := s.Database()

	if databaseHealth.Status == Unhealthy {
		overallStatus = Unhealthy
	}

	return &SystemState{
		Status:   overallStatus,
		Database: databaseHealth.Status,
	}
}

type SystemState struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}
