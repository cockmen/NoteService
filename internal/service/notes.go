package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
		c.Set("email", claims.Email)

		return next(c)
	}
}

func (s *Service) QOTD() (string, error) {
	resp, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		s.logger.Error("can`t get the quote")
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("can`t read the quote")
		return "", err
	}
	var quote struct {
		Quote struct {
			Body   string `json:"body"`
			Author string `json:"author"`
		} `json:"quote"`
	}

	if err := json.Unmarshal(body, &quote); err != nil {
		s.logger.Error("can`t unmarshall the quote")
		return "", err
	}
	return fmt.Sprintf("%s - %s", quote.Quote.Body, quote.Quote.Author), nil
}

func (s *Service) GetNotes(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		s.logger.Error("Unauthorized")
		return c.JSON(s.NewError(Unauthorized))
	}

	repo := s.notesRepo

	notes, err := repo.RGetNotes(email)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: notes})
}

func (s *Service) GetNoteById(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		s.logger.Error("Unathorized")
		return c.JSON(s.NewError(Unauthorized))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.notesRepo

	noteOwner, err := repo.RGetNoteOwner(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	if noteOwner != email {
		s.logger.Error(err)
		return c.JSON(s.NewError(Unauthorized))
	}

	note, err := repo.RGetNoteById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: note})
}

func (s *Service) CreateNewNote(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		s.logger.Error("trouble with email")
		return c.JSON(s.NewError(Unauthorized))
	}
	var note Note
	err := c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	quote, err := s.QOTD()
	if err != nil {
		s.logger.Error("couldn`t generate qoute")
		return c.JSON(s.NewError(InternalServerError))
	}

	repo := s.notesRepo

	note.UserEmail = email

	err = repo.RCreateNewNote(note.Title, note.Body, email)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"note":  note,
		"quote": quote,
	})
}

func (s *Service) DeleteNote(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		s.logger.Error("Unathorized")
		return c.JSON(s.NewError(Unauthorized))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.notesRepo

	noteOwner, err := repo.RGetNoteOwner(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidRequest))
	}
	if noteOwner != email {
		s.logger.Error(err)
		return c.JSON(s.NewError(Unauthorized))
	}

	err = repo.RDeleteNote(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (s *Service) UpdateNote(c echo.Context) error {
	email, ok := c.Get("email").(string)
	if !ok {
		s.logger.Error("Unathorized")
		return c.JSON(s.NewError(Unauthorized))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	var note Note
	if err := c.Bind(&note); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	repo := s.notesRepo

	noteOwner, err := repo.RGetNoteOwner(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidRequest))
	}
	if noteOwner != email {
		s.logger.Error(err)
		return c.JSON(s.NewError(Unauthorized))
	}

	err = repo.RUpdateNote(note.Title, note.Body, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Updated")
}
