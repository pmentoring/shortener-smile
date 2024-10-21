package routing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"shortener-smile/internal/common"
	appactions "shortener-smile/internal/common/http_actions"
	"shortener-smile/internal/shortener/http_actions"
)

func Register(engine *gin.Engine, db *sql.DB, ctx *common.ApplicationContext) {
	shortenURLAction := http_actions.NewShortenUrlAction(db, ctx)
	unshortenURLAction := http_actions.NewRedirectFromShortenAction(db)

	engine.POST("/shorten", shortenURLAction.UrlShorten)
	engine.GET("/:urlCode", unshortenURLAction.Redirect)
	engine.GET("/healthcheck", appactions.HandleHealth)
}
