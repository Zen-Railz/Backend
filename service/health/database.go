package health

func (s *Service) Database() *DatabaseState {
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

type DatabaseState struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
