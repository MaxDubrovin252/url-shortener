package service

import "url-shortener/pkg/repository"

type URL interface {
	Create(url string, alias string) (int64, error)
	Get(alias string) (string, error)
	Delete(alias string) error
}

type Service struct {
	URL
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URL: NewURLService(repos),
	}
}
