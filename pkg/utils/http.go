package utils

import (
	"io"
	"net/http"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

type HttpGetResponse struct {
	Body  []byte
	Error error
}

func Get(url string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, wg *sync.WaitGroup, ch chan HttpGetResponse) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		MessageWithReply(s, m, "Error fetching data", logger)
		logger.Error().Err(err).Msg("Error fetching data")
		ch <- HttpGetResponse{
			Body:  nil,
			Error: err,
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		MessageWithReply(s, m, "Error reading data", logger)
		logger.Error().Err(err).Msg("Error reading data")
		ch <- HttpGetResponse{
			Body:  nil,
			Error: err,
		}
	}

	ch <- HttpGetResponse{
		Body:  body,
		Error: nil,
	}
}
