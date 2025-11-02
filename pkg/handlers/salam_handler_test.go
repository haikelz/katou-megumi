package handlers_test

import (
	"katou-megumi/pkg/handlers"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestSalamHandler_BotMessage(t *testing.T) {
	// Setup - bot sends message to itself
	logger := zerolog.Nop()
	session := &discordgo.Session{}
	session.State = &discordgo.State{}
	session.State.User = &discordgo.User{
		ID: "bot-user-id",
	}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "bot-user-id", // Same as bot ID - should return early
	}

	// Execute - should return early without sending message
	assert.NotPanics(t, func() {
		handlers.SalamHandler(session, message, &logger, "")
	})
}

func TestSalamHandler_UserMessage(t *testing.T) {
	// Setup - user sends message
	logger := zerolog.Nop()
	session := &discordgo.Session{}
	session.State = &discordgo.State{}
	session.State.User = &discordgo.User{
		ID: "bot-user-id",
	}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "user-id", // Different from bot ID
	}

	// Execute - should try to send "Assalamu'alaikum" message
	assert.NotPanics(t, func() {
		handlers.SalamHandler(session, message, &logger, "")
	})
}

func TestSalamHandler_WithCommand(t *testing.T) {
	// Setup
	logger := zerolog.Nop()
	session := &discordgo.Session{}
	session.State = &discordgo.State{}
	session.State.User = &discordgo.User{
		ID: "bot-user-id",
	}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "user-id",
	}

	// Execute with command parameter
	assert.NotPanics(t, func() {
		handlers.SalamHandler(session, message, &logger, "test-command")
	})
}
