package handlers

import (
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func PingHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	utils.MessageWithReply(s, m, "Pong!", logger)
}
