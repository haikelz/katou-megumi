package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

// An utility function for discord's sending message with zap logger

// Reply user's message
func MessageWithReply(s *discordgo.Session, m *discordgo.MessageCreate, content string, logger *zerolog.Logger) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, content, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("Error sending message")
		return
	}
}

// Send message without mentioned the user
func Message(s *discordgo.Session, m *discordgo.MessageCreate, content string, logger *zerolog.Logger) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		logger.Error().Err(err).Msg("Error sending message")
		return
	}
}

// Send message with embed reply. Suitable for sending image or video.
func MessageWithEmbedReply(s *discordgo.Session, m *discordgo.MessageCreate, embed *discordgo.MessageEmbed, logger *zerolog.Logger) {
	_, err := s.ChannelMessageSendEmbedReply(m.ChannelID, embed, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("Error sending message")
		return
	}
}
