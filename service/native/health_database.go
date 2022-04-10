package native

type DatabaseState struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
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
