package http_actions

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-smile/internal/common"
	"shortener-smile/internal/shortener/repository"
	"shortener-smile/internal/shortener/service"
)

type ShortenURLAction struct {
	shortenLinkService *service.ShortenLinkService
}

type ShortenUrlRequest struct {
	URL   string `json:"url" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func NewShortenUrlAction(db *sql.DB, ctx *common.ApplicationContext) *ShortenURLAction {
	return &ShortenURLAction{
		shortenLinkService: service.NewShortenLinkService(repository.NewLinksRepository(db), ctx),
	}
}

func (a ShortenURLAction) UrlShorten(ctx *gin.Context) {
	var shortenUrlRequest ShortenUrlRequest

	if err := ctx.ShouldBind(&shortenUrlRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	link, err := a.shortenLinkService.CreateShortenLink(shortenUrlRequest.Title, shortenUrlRequest.URL)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": "error while saving shorten url"})
		return
	}

	ctx.JSON(200, gin.H{"url": link.ShortenLink})
}
