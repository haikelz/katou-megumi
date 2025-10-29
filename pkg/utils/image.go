package utils

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func ImageUrlToBase64(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, imageUrl string) string {
	client := &http.Client{}

	image, err := client.Get(imageUrl)
	if err != nil {
		MessageWithReply(s, m, "Error getting image", logger)
	}
	defer image.Body.Close()

	imageBytes, err := io.ReadAll(image.Body)
	if err != nil {
		MessageWithReply(s, m, "Error reading image", logger)
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
	return base64Image
}
