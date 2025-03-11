package service

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func (s *Service) JWTCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) == 0 {
			s.logger.Error("missing token")
			return c.JSON(s.NewError(MissingToken))
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			s.logger.Error(authHeader)
			return c.JSON(s.NewError(InvalidParams))
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			s.logger.Error(err)
			return c.JSON(s.NewError(InvalidToken))
		}
		c.Set("id", claims.Id)

		return next(c)
	}
}
