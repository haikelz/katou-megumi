package handlers

import (
	"encoding/json"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func JokeHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	imageUrl, err := getJokeImage(s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke image", logger)
		logger.Error().Err(err).Msg("Error getting joke image")
		return
	}

	jokeText, err := getJokeText(s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke text", logger)
		logger.Error().Err(err).Msg("Error getting joke text")
		return
	}

	utils.MessageWithEmbedReply(s, m, &discordgo.MessageEmbed{
		Title: jokeText,
		Image: &discordgo.MessageEmbedImage{URL: imageUrl},
	}, logger)
}

func getJokeImage(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) (string, error) {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)
	go utils.Get(utils.Env().JOKES_API_URL+"/api/image/random", s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		logger.Error().Err(response.Error).Msg("Error fetching jokes")
		return "", response.Error
	}

	var jokeImageResponse entities.JokeImageResponse
	err := json.Unmarshal(response.Body, &jokeImageResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling jokes", logger)
		logger.Error().Err(err).Msg("Error unmarshalling jokes")
		return "", err
	}

	return jokeImageResponse.Data.Url, nil
}

func getJokeText(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) (string, error) {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)
	go utils.Get(utils.Env().JOKES_API_URL+"/api/text/random", s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		logger.Error().Err(response.Error).Msg("Error fetching jokes")
		return "", response.Error
	}

	var jokeTextResponse entities.JokeTextResponse
	err := json.Unmarshal(response.Body, &jokeTextResponse)
	if err != nil {
		logger.Error().Err(err).Msg("Error unmarshalling jokes")
		return "", err
	}

	return jokeTextResponse.Data, nil
}
