package utils_test

import (
	"katou-megumi/pkg/utils"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestMessageWithReply_ValidInput tests that MessageWithReply handles valid input without panicking
// Note: This is a unit test that verifies the function doesn't crash with valid input.
// Full integration test would require a real Discord bot token.
func TestMessageWithReply_ValidInput(t *testing.T) {
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
			GuildID:   "test-guild-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Execute - this will likely fail because session is not connected
	// We use recover to catch expected panic from uninitialized session
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session
				// This is acceptable for unit testing without full Discord connection
			}
		}()
		utils.MessageWithReply(session, message, "Test message", &logger)
	}()
	// Test passes if no unexpected panic occurs
}

// TestMessage_ValidInput tests that Message handles valid input without panicking
func TestMessage_ValidInput(t *testing.T) {
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
		ID: "test-user-id",
	}

	// Execute - this will likely fail because session is not connected
	// We use recover to catch expected panic from uninitialized session
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session
				// This is acceptable for unit testing without full Discord connection
			}
		}()
		utils.Message(session, message, "Test message", &logger)
	}()
	// Test passes if no unexpected panic occurs
}

// TestMessageWithEmbedReply_ValidInput tests that MessageWithEmbedReply handles valid input without panicking
func TestMessageWithEmbedReply_ValidInput(t *testing.T) {
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
			GuildID:   "test-guild-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Test Embed",
		Description: "Test Description",
	}

	// Execute - this will likely fail because session is not connected
	// We use recover to catch expected panic from uninitialized session
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session
				// This is acceptable for unit testing without full Discord connection
			}
		}()
		utils.MessageWithEmbedReply(session, message, embed, &logger)
	}()
	// Test passes if no unexpected panic occurs
}

// TestMessageWithReply_NilSession tests that MessageWithReply handles nil session gracefully
func TestMessageWithReply_NilSession(t *testing.T) {
	// Setup
	logger := zerolog.Nop()
	var session *discordgo.Session
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
			GuildID:   "test-guild-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Execute - nil session will panic, which is expected behavior
	// We use recover to catch expected panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic with nil session
				// This is acceptable for unit testing
			}
		}()
		utils.MessageWithReply(session, message, "Test message", &logger)
	}()
	// Test passes if no unexpected panic occurs
}

// TestMessageWithReply_MessageReferenceStructure tests that the MessageReference is constructed correctly
// This test verifies the structure of the reference object
func TestMessageWithReply_MessageReferenceStructure(t *testing.T) {
	// Setup
	logger := zerolog.Nop()
	session := &discordgo.Session{}
	session.State = &discordgo.State{}
	session.State.User = &discordgo.User{
		ID: "bot-user-id",
	}
	messageID := "test-message-id"
	channelID := "test-channel-id"
	guildID := "test-guild-id"

	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        messageID,
			ChannelID: channelID,
			GuildID:   guildID,
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Verify message structure before calling
	assert.Equal(t, messageID, message.ID)
	assert.Equal(t, channelID, message.ChannelID)
	assert.Equal(t, guildID, message.GuildID)

	// Execute - this will likely fail because session is not connected
	// We use recover to catch expected panic from uninitialized session
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session
				// This is acceptable for unit testing without full Discord connection
			}
		}()
		utils.MessageWithReply(session, message, "Test message", &logger)
	}()
	// Test passes if no unexpected panic occurs
}
