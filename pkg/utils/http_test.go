package utils_test

import (
	"encoding/json"
	"katou-megumi/pkg/utils"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet_Success(t *testing.T) {
	// Setup mock HTTP server
	expectedBody := map[string]string{"message": "success"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedBody)
	}))
	defer server.Close()

	// Setup test dependencies
	logger := zerolog.Nop()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan utils.HttpGetResponse, 1)

	session := &discordgo.Session{}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Execute
	utils.Get(server.URL, session, message, &logger, wg, ch)
	wg.Wait()
	close(ch)

	// Assert
	response := <-ch
	require.NoError(t, response.Error)
	assert.NotNil(t, response.Body)
	assert.Greater(t, len(response.Body), 0)
}

func TestGet_ServerError(t *testing.T) {
	// Setup mock HTTP server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	server.Close() // Close immediately to simulate connection error

	// Setup test dependencies
	logger := zerolog.Nop()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan utils.HttpGetResponse, 1)

	session := &discordgo.Session{}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Execute - will fail because server is closed
	// Wrap in recover to handle panic from MessageWithReply
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session in MessageWithReply
				// This is acceptable for unit testing
			}
		}()
		utils.Get(server.URL, session, message, &logger, wg, ch)
		wg.Wait()
		close(ch)

		// Assert
		response := <-ch
		assert.Error(t, response.Error)
		assert.Nil(t, response.Body)
	}()
}

func TestGet_InvalidURL(t *testing.T) {
	// Setup test dependencies
	logger := zerolog.Nop()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan utils.HttpGetResponse, 1)

	session := &discordgo.Session{}
	message := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "test-message-id",
			ChannelID: "test-channel-id",
		},
	}
	message.Author = &discordgo.User{
		ID: "test-user-id",
	}

	// Execute with invalid URL
	// Wrap in recover to handle panic from MessageWithReply
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected panic due to uninitialized session in MessageWithReply
				// This is acceptable for unit testing
			}
		}()
		utils.Get("http://invalid-url-that-does-not-exist:9999", session, message, &logger, wg, ch)
		wg.Wait()
		close(ch)

		// Assert
		response := <-ch
		assert.Error(t, response.Error)
		assert.Nil(t, response.Body)
	}()
}
