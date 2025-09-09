package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"katou-megumi/pkg/configs"
	"katou-megumi/pkg/utils"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"google.golang.org/genai"
)

func GeminiHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	client := configs.NewGemini(context.Background(), utils.Env().GEMINI_API_KEY)

	// handle if user also attach an image to the message
	if len(m.Attachments) > 0 {
		base64Image := utils.ImageUrlToBase64(s, m, logger, m.Attachments[0].URL)

		imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			utils.MessageWithReply(s, m, "Error decoding image", logger)
			logger.Error("Error decoding image", zap.Error(err))
			return
		}

		imgPart := &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: m.Attachments[0].ContentType,
				Data:     imageBytes,
			},
		}

		response, err := client.Models.GenerateContent(context.Background(), utils.GEMINI_MODEL, []*genai.Content{
			{
				Role: utils.GEMINI_ROLE,
				Parts: []*genai.Part{
					genai.NewPartFromText("Jawab pertanyaan ini dengan bahasa Indonesia: " + command),
					imgPart,
				},
			},
		}, &genai.GenerateContentConfig{
			ResponseMIMEType: "text/plain",
		})
		if err != nil {
			utils.MessageWithReply(s, m, "Error generating content", logger)
			logger.Error("Error generating content", zap.Error(err))
			return
		}

		if response.Text() == "" {
			utils.MessageWithReply(s, m, "Error generating content", logger)
			logger.Error("Error generating content", zap.Error(errors.New("error generating content")))
			return
		}

		utils.MessageWithReply(s, m, response.Text(), logger)
	} else {
		response, err := client.Models.GenerateContent(context.Background(), utils.GEMINI_MODEL, []*genai.Content{
			{
				Role: utils.GEMINI_ROLE,
				Parts: []*genai.Part{
					{
						Text: "Jawab pertanyaan ini dengan bahasa Indonesia: " + command,
					},
				},
			},
		}, &genai.GenerateContentConfig{
			ResponseMIMEType: "text/plain",
		})
		if err != nil {
			utils.MessageWithReply(s, m, "Error generating content", logger)
			logger.Error("Error generating content", zap.Error(err))
			return
		}

		if response.Text() == "" {
			utils.MessageWithReply(s, m, "Error generating content", logger)
			logger.Error("Error generating content", zap.Error(errors.New("error generating content")))
			return
		}

		utils.MessageWithReply(s, m, response.Text(), logger)
	}
}
