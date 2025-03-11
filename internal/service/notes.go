package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetNotes(c echo.Context) error {
	id, ok := c.Get("id").(int)
	if !ok {
		s.logger.Error("Unauthorized")
		return c.JSON(s.NewError(Unauthorized))
	}

	repo := s.notesRepo

	notes, err := repo.RGetNotes(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: notes})
}

func (s *Service) GetNoteById(c echo.Context) error {
	_, ok := c.Get("id").(int)
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

	note, err := repo.RGetNoteById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: note})
}

func (s *Service) CreateNewNote(c echo.Context) error {
	id, ok := c.Get("id").(int)
	if !ok {
		s.logger.Error("Unathorized")
		return c.JSON(s.NewError(Unauthorized))
	}
	var note Note
	err := c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	quote, err := s.QuoteOfTheDay()
	if err != nil {
		s.logger.Error("couldn`t generate qoute")
		return c.JSON(s.NewError(InternalServerError))
	}

	repo := s.notesRepo

	note.UserId = id

	err = repo.RCreateNewNote(note.Title, note.Body, id)
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
	_, ok := c.Get("id").(int)
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

	err = repo.RDeleteNote(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (s *Service) UpdateNote(c echo.Context) error {
	_, ok := c.Get("id").(int)
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

	err = repo.RUpdateNote(note.Title, note.Body, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Updated")
}
