package native

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

func (s *Service) DatabaseHealth() *DatabaseState {
	healthErr := s.databaseRepo.Ping()
	if healthErr != nil {
		healthErr.Trace()
		s.logger.Error(healthErr.Elaborate(), nil)
		return &DatabaseState{
			Status:  Unhealthy,
			Message: healthErr.Elaborate(),
		}
	}

	return &DatabaseState{
		Status: Healthy,
	}
}
