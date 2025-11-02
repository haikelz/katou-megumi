package handlers_test

import (
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/handlers"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestJadwalSholatHandler_EmptyCommand(t *testing.T) {
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
		ID:       "user-id",
		Username: "testuser",
	}

	// Execute with empty command
	assert.NotPanics(t, func() {
		handlers.JadwalSholatHandler(session, message, &logger, "")
	})
}

func TestJadwalSholatHandler_ShortCommand(t *testing.T) {
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
		ID:       "user-id",
		Username: "testuser",
	}

	// Execute with short command (length <= 2)
	assert.NotPanics(t, func() {
		handlers.JadwalSholatHandler(session, message, &logger, "ab")
	})
}

// Test JadwalSholatHandler structure validation
func TestJadwalSholatHandler_StructureValidation(t *testing.T) {
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
		ID:       "user-id",
		Username: "testuser",
	}

	// Test various command lengths
	testCases := []struct {
		name    string
		command string
	}{
		{"Empty command", ""},
		{"Single character", "a"},
		{"Two characters", "ab"},
		{"Valid city name", "Jakarta"},
		{"Long city name", "Jakarta Pusat"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				handlers.JadwalSholatHandler(session, message, &logger, tc.command)
			})
		})
	}
}

// Test helper to validate JadwalSholatResponse structure
func TestJadwalSholatResponse_Structure(t *testing.T) {
	// Setup sample response
	today := time.Now().Format("2006-01-02")
	response := entities.JadwalSholatResponse{
		Data: &entities.JadwalSholat{
			Daerah: "Jakarta",
			Lokasi: "DKI Jakarta",
			Jadwal: &entities.Jadwal{
				Imsak:   "04:30",
				Subuh:   "04:40",
				Terbit:  "05:50",
				Dhuha:   "06:20",
				Dzuhur:  "12:00",
				Ashar:   "15:30",
				Maghrib: "18:00",
				Isya:    "19:15",
				Date:    today,
			},
		},
	}

	// Assert structure
	assert.NotNil(t, response.Data)
	assert.Equal(t, "Jakarta", response.Data.Daerah)
	assert.Equal(t, "DKI Jakarta", response.Data.Lokasi)
	assert.NotNil(t, response.Data.Jadwal)
	assert.Equal(t, "04:30", response.Data.Jadwal.Imsak)
	assert.Equal(t, "04:40", response.Data.Jadwal.Subuh)
	assert.Equal(t, "05:50", response.Data.Jadwal.Terbit)
	assert.Equal(t, "06:20", response.Data.Jadwal.Dhuha)
	assert.Equal(t, "12:00", response.Data.Jadwal.Dzuhur)
	assert.Equal(t, "15:30", response.Data.Jadwal.Ashar)
	assert.Equal(t, "18:00", response.Data.Jadwal.Maghrib)
	assert.Equal(t, "19:15", response.Data.Jadwal.Isya)
}

// Test JadwalSholaCityIdResponse structure
func TestJadwalSholaCityIdResponse_Structure(t *testing.T) {
	// Setup sample response
	response := entities.JadwalSholaCityIdResponse{
		Data: []*entities.JadwalSholaCityIdResponseData{
			{
				Id:     "123",
				Lokasi: "Jakarta",
			},
			{
				Id:     "456",
				Lokasi: "Bandung",
			},
		},
	}

	// Assert structure
	assert.NotNil(t, response.Data)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, "123", response.Data[0].Id)
	assert.Equal(t, "Jakarta", response.Data[0].Lokasi)
	assert.Equal(t, "456", response.Data[1].Id)
	assert.Equal(t, "Bandung", response.Data[1].Lokasi)
}
