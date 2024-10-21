package service

import (
	"shortener-smile/internal/shortener/domain/models"
	"shortener-smile/internal/shortener/repository"
)

type UnshortenLinkService struct {
	repo repository.LinksRepository
}

func NewUnshortenLinkService(repository repository.LinksRepository) *UnshortenLinkService {
	return &UnshortenLinkService{
		repo: repository,
	}
}

func (s UnshortenLinkService) Unshorten(linkCode string) (*models.Link, error) {
	link, err := s.repo.FindLinkByCode(linkCode)
	if err != nil {
		return nil, err
	}

	go s.IncrementFollowsCounter(link)

	return link, nil
}

func (s UnshortenLinkService) IncrementFollowsCounter(link *models.Link) {
	err := s.repo.IncrementFollowsCounter(link)
	if err != nil {
		return
	}
}
