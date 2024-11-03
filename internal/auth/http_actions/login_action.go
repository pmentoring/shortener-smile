package http_actions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-smile/internal/auth/http_actions/request"
	"shortener-smile/internal/auth/service"
)

type LoginAction struct {
	loginService *service.LoginUserService
}

func NewLoginAction(s *service.LoginUserService) *LoginAction {
	return &LoginAction{s}
}

func (action *LoginAction) LoginUser(ctx *gin.Context) {
	var loginRequest request.LoginUserRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token, err := action.loginService.CreateJwtByLoginAndPassword(loginRequest.Login, loginRequest.Password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
