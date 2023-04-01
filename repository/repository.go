package repository

import (
	"errors"

	"go-theapi/model"
	"go-theapi/token"
	"go-theapi/util"

	"github.com/labstack/echo/v4"
)

func GenerateToken(c echo.Context, user string) (*model.M_Response_session, error) {
	var session model.M_Response_session

	config, _ := util.LoadConfig(".")
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		err := errors.New("cannot create token maker")
		return &session, err
	}

	accessToken, accessPayload, err := tokenMaker.CreateToken(
		user,
		config.AccessTokenDuration,
	)
	if err != nil {
		err := errors.New("cannot generate access token")
		return &session, err
	}

	result := &model.M_Response_session{
		SessionID:        accessPayload.ID,
		Username:         user,
		AccessToken:      accessToken,
		Access_ExpiresAt: accessPayload.ExpiredAt,
	}

	return result, nil
}
