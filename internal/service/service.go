package service

import(
	"github.com/VadimShara/coursework_db_shop/internal/repo"
)

type Service struct {
	db         *repo.DB
}

func (s *Service) NewService(db *repo.DB) *Service {
	return &Service{
		db: db,
	}
}