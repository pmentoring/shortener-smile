package http_actions

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-smile/internal/shortener/repository"
	"shortener-smile/internal/shortener/service"
)

type RedirectFromShortenAction struct {
	service *service.UnshortenLinkService
}

func NewRedirectFromShortenAction(db *sql.DB) *RedirectFromShortenAction {
	return &RedirectFromShortenAction{
		service: service.NewUnshortenLinkService(repository.NewLinksRepository(db)),
	}
}

func (a *RedirectFromShortenAction) Redirect(ctx *gin.Context) {
	param := ctx.Param("urlCode")

	link, err := a.service.Unshorten(param)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("redirect:", link)

	ctx.Redirect(http.StatusPermanentRedirect, link.FullLink)
}
