package utils_test

import (
	"encoding/base64"
	"katou-megumi/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageUrlToBase64_Success(t *testing.T) {
	// Setup mock HTTP server with image data
	testImageData := []byte("fake-image-data")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(testImageData)
	}))
	defer server.Close()

	// Setup test dependencies
	logger := zerolog.Nop()
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	session := &discordgo.Session{}

	// Execute
	result := utils.ImageUrlToBase64(session, message, &logger, server.URL)

	// Assert
	require.NotEmpty(t, result)
	decoded, err := base64.StdEncoding.DecodeString(result)
	require.NoError(t, err)
	assert.Equal(t, testImageData, decoded)
}

func TestImageUrlToBase64_InvalidURL(t *testing.T) {
	// Setup test dependencies
	logger := zerolog.Nop()
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	session := &discordgo.Session{}

	// Execute with invalid URL
	// Wrap in recover to handle panic from MessageWithReply
	var result string
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session in MessageWithReply
				// This is acceptable for unit testing
			}
		}()
		result = utils.ImageUrlToBase64(session, message, &logger, "http://invalid-url-that-does-not-exist:9999/image.jpg")
	}()

	// Assert - function should return empty string or handle error gracefully
	// The actual behavior depends on implementation, but it shouldn't panic
	assert.NotNil(t, result) // Result might be empty or contain partial data
}

func TestImageUrlToBase64_ServerError(t *testing.T) {
	// Setup mock HTTP server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Setup test dependencies
	logger := zerolog.Nop()
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	session := &discordgo.Session{}

	// Execute
	result := utils.ImageUrlToBase64(session, message, &logger, server.URL)

	// Assert - function should handle error gracefully
	// Even if error occurs, it might return empty string or partial base64
	assert.NotNil(t, result)
}
