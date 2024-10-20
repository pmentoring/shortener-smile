package routing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"shortener-smile/internal/common"
	appactions "shortener-smile/internal/common/http_actions"
	"shortener-smile/internal/shortener/http_actions"
)

func Register(engine *gin.Engine, db *sql.DB, ctx *common.ApplicationContext) {
	shortenURLAction := http_actions.New(db, ctx)

	engine.POST("/shorten", shortenURLAction.UrlShorten)

	engine.GET("/healthcheck", appactions.HandleHealth)
}
