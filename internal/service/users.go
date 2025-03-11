package service

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secretkey")

// Регистрация пользователя
func (s *Service) Registration(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidRequest))
	}
	if _, err := s.usersRepo.GetUserByEmail(user.Email); err == nil {
		s.logger.Error("user already exist")
		return c.JSON(s.NewError(UserExist))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(ErrWithPass))
	}
	repo := s.usersRepo
	err = repo.CreateUser(user.Email, string(hashedPassword))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "Created")
}

// Аутентификация пользователя с выдачей jwt-токена
func (s *Service) Login(c echo.Context) error {
	repo := s.usersRepo

	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidRequest))
	}

	user, err := repo.GetUserByEmail(req.Email)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(UserNotFound))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidPassword))
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email

	t, err := token.SignedString(jwtKey)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(ErrWithPass))
	}

	return c.JSON(http.StatusOK, Response{Object: t})
}
