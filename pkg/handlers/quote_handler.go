package handlers

import (
	"encoding/json"
	"fmt"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func QuoteHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if command == "" {
		quotes := GetQuotes(s, m, logger)
		_, err := s.ChannelMessageSendReply(m.ChannelID, quotes, &discordgo.MessageReference{
			MessageID: m.ID,
			ChannelID: m.ChannelID,
			GuildID:   m.GuildID,
		})
		if err != nil {
			logger.Error("Error sending message", zap.Error(err))
			return
		}
		return
	}

	quote := GetQuote(command, s, m, logger)

	utils.MessageWithReply(s, m, quote, logger)
}

func GetQuote(anime string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) string {
	body := utils.Get(utils.Env().ANIME_QUOTE_API_URL+"/api/getbyanime?anime="+anime+"&page=1", s, m, logger)

	var quoteResponse entities.QuoteResponse
	err := json.Unmarshal(body, &quoteResponse)
	if err != nil {
		logger.Error("Error unmarshalling quote", zap.Error(err))
		return ""
	}

	content := fmt.Sprintf("%s - %s - %s - %s\n\n", quoteResponse.Result[0].English, quoteResponse.Result[0].Indo, quoteResponse.Result[0].Anime, quoteResponse.Result[0].Character)
	return content
}

func GetQuotes(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) string {
	content := ""

	body := utils.Get(utils.Env().ANIME_QUOTE_API_URL+"/api/getrandom", s, m, logger)

	var quoteResponse entities.QuotesResponse
	err := json.Unmarshal(body, &quoteResponse)
	if err != nil {
		logger.Error("Error unmarshalling quote", zap.Error(err))
		return ""
	}

	for _, v := range quoteResponse.Result {
		content += fmt.Sprintf("%s - %s - %s - %s\n\n", v.English, v.Indo, v.Anime, v.Character)
	}

	return content
}
