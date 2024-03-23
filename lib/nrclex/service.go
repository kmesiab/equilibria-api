package nrclex

import "time"

// Service provides services related to NrcLex entities.
type Service struct {
	repo *Repository
}

// NewService creates a new Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateNrcLex creates a new NrcLex.
func (service *Service) CreateNrcLex(nrcLex *NrcLex) error {
	return service.repo.Create(nrcLex)
}

// FindNrcLexByID retrieves a NrcLex by its ID.
func (service *Service) FindNrcLexByID(id int64) (*NrcLex, error) {
	return service.repo.FindByID(id)
}

// UpdateNrcLex updates an existing NrcLex.
func (service *Service) UpdateNrcLex(nrcLex *NrcLex) error {
	return service.repo.Update(nrcLex)
}

// DeleteNrcLex deletes a NrcLex.
func (service *Service) DeleteNrcLex(id int64) error {
	return service.repo.Delete(id)
}

func (service *Service) FindRangeByUserID(userID int64, limit, offset int, startDate, endDate time.Time) (*[]NrcLex, error) {

	return service.repo.FindRangeByUserID(userID, limit, offset, startDate, endDate)
}
