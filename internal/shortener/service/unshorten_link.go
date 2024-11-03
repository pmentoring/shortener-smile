package service

import (
	"log/slog"
	"net/url"
	"shortener-smile/internal/shortener/domain/models"
	"shortener-smile/internal/shortener/repository"
	"strings"
)

type UrlShortenerService struct {
	repo repository.LinksRepository
	log  *slog.Logger
}

func NewUnshortenLinkService(repository repository.LinksRepository, log *slog.Logger) *UrlShortenerService {
	return &UrlShortenerService{
		repo: repository,
		log:  log,
	}
}

func (s UrlShortenerService) Unshorten(linkCode string) (*models.Link, error) {
	link, err := s.repo.FindLinkByCode(linkCode)
	if err != nil {
		s.log.Info(err.Error())
		return nil, err
	}

	go s.incrementFollowsCounter(link)

	return link, nil
}

func (s UrlShortenerService) GetLinkByCode(link string) (*models.Link, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		s.log.Info(err.Error())
		return nil, err
	}

	path := strings.TrimPrefix(parsedURL.Path, "/")

	linkRecord, err := s.repo.FindLinkByCode(path)
	if err != nil {
		s.log.Info(err.Error())
		return nil, err
	}

	go s.incrementFollowsCounter(linkRecord)

	return linkRecord, nil
}

func (s UrlShortenerService) incrementFollowsCounter(link *models.Link) {
	err := s.repo.IncrementFollowsCounter(link)
	if err != nil {
		s.log.Info(err.Error())
		return
	}
}
