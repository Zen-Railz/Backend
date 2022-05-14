package railway

import "zenrailz/errorr"

func (s *Service) Lines() ([]Line, errorr.Entity) {
	dbLines, dbResErr := s.railwayRepo.Lines()
	if dbResErr != nil {
		dbResErr.Trace()
		s.logger.Error(dbResErr.Elaborate(), nil)
		return nil, dbResErr
	}

	lines := []Line{}
	for _, dbLine := range dbLines {
		lines = append(lines, Line{
			Name:         dbLine.Name,
			Code:         dbLine.Code,
			Type:         dbLine.Type,
			IsActive:     dbLine.IsActive,
			Announcement: dbLine.Announcement,
		})
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
