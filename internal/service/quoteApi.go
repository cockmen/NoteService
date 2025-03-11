package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var quote struct {
	Quote struct {
		Body   string `json:"body"`
		Author string `json:"author"`
	} `json:"quote"`
}

func (s *Service) QuoteOfTheDay() (string, error) {
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

	if err := json.Unmarshal(body, &quote); err != nil {
		s.logger.Error("can`t unmarshall the quote")
		return "", err
	}
	return fmt.Sprintf("%s - %s", quote.Quote.Body, quote.Quote.Author), nil
}
