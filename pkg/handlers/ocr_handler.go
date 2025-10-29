package handlers

import (
	"io"
	"katou-megumi/pkg/utils"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/otiai10/gosseract/v2"
	"github.com/rs/zerolog"
)

func OcrHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	client := &http.Client{}
	imageUrl := m.Attachments[0].URL

	ocr := gosseract.NewClient()
	defer ocr.Close()

	image, err := client.Get(imageUrl)
	if err != nil {
		utils.MessageWithReply(s, m, "Error getting image", logger)
		logger.Error().Err(err).Msg("Error getting image")
		return
	}
	defer image.Body.Close()

	imageBytes, err := io.ReadAll(image.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading image", logger)
		logger.Error().Err(err).Msg("Error reading image")
		return
	}

	ocr.SetImageFromBytes(imageBytes)
	text, _ := ocr.Text()

	utils.MessageWithReply(s, m, text, logger)
}
