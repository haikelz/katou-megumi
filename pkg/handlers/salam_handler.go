package handlers

import (
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func SalamHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	utils.MessageWithReply(s, m, "Assalamu'alaikum", logger)
}
