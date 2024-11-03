package service

import (
	"fmt"
	"github.com/deatil/go-encoding/encoding"
	"log/slog"
	"shortener-smile/internal/common"
	"shortener-smile/internal/shortener/domain/models"
	"shortener-smile/internal/shortener/repository"
	"strconv"
)

type ShortenLinkService struct {
	repo repository.LinksRepository
	ctx  *common.ApplicationContext
	log  *slog.Logger
}

func NewShortenLinkService(repo repository.LinksRepository, ctx *common.ApplicationContext, log *slog.Logger) *ShortenLinkService {
	return &ShortenLinkService{
		repo: repo,
		ctx:  ctx,
		log:  log,
	}
}

func (sl ShortenLinkService) CreateShortenLink(title string, url string) (*models.Link, error) {
	nextId, err := sl.repo.GetNextId()
	if err != nil {
		return nil, err
	}

	shortenUrlCode := encoding.
		FromString(sl.ctx.InstanceId + "_" + strconv.Itoa(nextId)).
		Base62Encode().
		ToString()

	fmt.Println(shortenUrlCode)

	link := &models.Link{
		Id:              nextId,
		Title:           title,
		FullLink:        url,
		ShortenLinkCode: shortenUrlCode,
		ShortenLink:     sl.ctx.AppBaseUrl + shortenUrlCode,
	}

	err = sl.repo.Save(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}
