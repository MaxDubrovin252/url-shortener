package service

import (
	"fmt"
	"url-shortener/pkg/repository"
)

type UrlService struct {
	repo *repository.Repository
}

func NewURLService(repo *repository.Repository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) Create(url string, alias string) (int64, error) {

	const operation = "service.url.create"
	if url == "" && alias == "" {

		return 0, fmt.Errorf("%s:%s", operation, ErrInvalidInput)
	}

	id, err := s.repo.Create(url, alias)

	if err != nil {
		return 0, fmt.Errorf("%s:%w", operation, err)
	}

	return id, nil

}

func (s *UrlService) Get(alias string) (string, error) {
	const operation = "service.url.get"

	if alias == "" {
		return "alias is empty", nil
	}

	url, err := s.repo.Get(alias)

	if err != nil {

		return "", fmt.Errorf("%s:%w", operation, err)
	}

	return url, nil
}
func (s *UrlService) Delete(alias string) error {
	const operation = "service.url.get"

	if alias == "" {
		return fmt.Errorf("%s:%w", operation, ErrInvalidInput)
	}

	if err := s.repo.Delete(alias); err != nil {
		return fmt.Errorf("%s:%w", operation, err)
	}

	return nil
}
