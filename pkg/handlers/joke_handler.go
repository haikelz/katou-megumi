package handlers

import (
	"encoding/json"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func JokeHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	imageUrl, err := getJokeImage(s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke image", logger)
		logger.Error("Error getting joke image", zap.Error(err))
		return
	}

	jokeText, err := getJokeText(s, m, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting joke text", logger)
		logger.Error("Error getting joke text", zap.Error(err))
		return
	}

	utils.MessageWithEmbedReply(s, m, &discordgo.MessageEmbed{
		Title: jokeText,
		Image: &discordgo.MessageEmbedImage{URL: imageUrl},
	}, logger)
}

func getJokeImage(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) (string, error) {
	body := utils.Get(utils.Env().JOKES_API_URL+"/api/image/random", s, m, logger)

	var jokeImageResponse entities.JokeImageResponse
	err := json.Unmarshal(body, &jokeImageResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling jokes", logger)
		logger.Error("Error unmarshalling jokes", zap.Error(err))
		return "", err
	}

	return jokeImageResponse.Data.Url, nil
}

func getJokeText(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) (string, error) {
	body := utils.Get(utils.Env().JOKES_API_URL+"/api/text/random", s, m, logger)

	var jokeTextResponse entities.JokeTextResponse
	err := json.Unmarshal(body, &jokeTextResponse)
	if err != nil {
		logger.Error("Error unmarshalling jokes", zap.Error(err))
		return "", err
	}

	return jokeTextResponse.Data, nil
}
