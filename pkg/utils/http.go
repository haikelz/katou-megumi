package utils

import (
	"io"
	"net/http"
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type HttpGetResponse struct {
	Body  []byte
	Error error
}

func Get(url string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, wg *sync.WaitGroup, ch chan HttpGetResponse) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		MessageWithReply(s, m, "Error fetching data", logger)
		logger.Error("Error fetching data", zap.Error(err))
		ch <- HttpGetResponse{
			Body:  nil,
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		MessageWithReply(s, m, "Error reading data", logger)
		logger.Error("Error reading data", zap.Error(err))
		ch <- HttpGetResponse{
			Body:  nil,
			Error: err,
		}
		return
	}

	ch <- HttpGetResponse{
		Body:  body,
		Error: nil,
	}
}
