package middlewares

import (
	"fmt"
	"go-theapi/token"
	"go-theapi/util"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func SetAuthMiddlewares(g *echo.Group) {
	g.Use(authMiddleware)
}

// AuthMiddleware creates a echo middleware for authorization
func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)

		config, _ := util.LoadConfig(".")
		tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if len(authorizationHeader) == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "authorization header is not provided")
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set(authorizationPayloadKey, payload)
		return next(c)

	}
}
