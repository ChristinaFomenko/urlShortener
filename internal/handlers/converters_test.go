package handlers

import (
	"github.com/ChristinaFomenko/shortener/internal/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToGetUrlsReply(t *testing.T) {
	tests := []struct {
		model []models.UserURL
		exp   []models.GetUrlsReply
	}{
		{
			model: []models.UserURL{
				{
					ShortURL:    "http://localhost:8080/abcde",
					OriginalURL: "https://yandex.ru",
				},
				{
					ShortURL:    "http://localhost:8080/qwerty",
					OriginalURL: "https://github.com",
				},
			},
			exp: []models.GetUrlsReply{
				{
					ShortURL:    "http://localhost:8080/abcde",
					OriginalURL: "https://yandex.ru",
				},
				{
					ShortURL:    "http://localhost:8080/qwerty",
					OriginalURL: "https://github.com",
				},
			},
		},
		{
			model: []models.UserURL{},
			exp:   []models.GetUrlsReply{},
		},
		{
			model: nil,
			exp:   []models.GetUrlsReply{},
		},
	}

	for _, tt := range tests {
		act := toGetUrlsReply(tt.model)

		assert.Equal(t, tt.exp, act)
	}
}
