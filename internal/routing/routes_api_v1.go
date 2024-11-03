package routing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	authactions "shortener-smile/internal/auth/http_actions"
	"shortener-smile/internal/auth/repository"
	"shortener-smile/internal/auth/service"
	"shortener-smile/internal/common"
	appactions "shortener-smile/internal/common/http_actions"
	shorteneractions "shortener-smile/internal/shortener/http_actions"
)

func Register(engine *gin.Engine, db *sql.DB, ctx *common.ApplicationContext) {
	shortenURLAction := shorteneractions.NewShortenUrlAction(db, ctx)
	unshortenURLAction := shorteneractions.NewRedirectFromShortenAction(db)

	loginUserAction := authactions.NewLoginAction(
		service.NewLoginUserService(
			repository.NewUserRepository(db),
			ctx,
			service.NewJWTService(ctx.SecretKey),
		),
	)

	engine.POST("/shorten", shortenURLAction.UrlShorten)
	engine.GET("/:urlCode", unshortenURLAction.Redirect)
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/auth/login", loginUserAction.LoginUser)
}
