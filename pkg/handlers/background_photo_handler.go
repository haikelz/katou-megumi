package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func BackgroundPhotoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	if command == "info" {
		utils.MessageWithReply(s, m, "Ini adalah perintah untuk mengubah warna background dari sebuah foto. Nama warna yang dimasukkan harus dalam bahasa Inggris.  Contoh: *!editphoto red*", logger)
		return
	}

	if command == "" {
		utils.MessageWithReply(s, m, "Mohon masukkan warna background yang diinginkan! Contoh: *!editphoto red*", logger)
		return
	}

	if len(m.Attachments) == 0 {
		utils.MessageWithReply(s, m, "Mohon sertakan gambar yang ingin diubah backgroundnya!", logger)
		return
	}

	client := &http.Client{}

	attachment := m.Attachments[0]
	if attachment.Width == 0 || attachment.Height == 0 {
		utils.MessageWithReply(s, m, "File yang dilampirkan bukan gambar yang valid!", logger)
		return
	}

	imageBase64 := utils.ImageUrlToBase64(s, m, logger, attachment.URL)

	removeBgData := entities.RemoveBgRequest{ImageFileB64: imageBase64, BgColor: command}

	jsonData, err := json.Marshal(&removeBgData)
	if err != nil {
		utils.MessageWithReply(s, m, "Error marshalling data", logger)
		logger.Error().Err(err).Msg("Error marshalling data")
		return
	}

	url := fmt.Sprintf("%s/v1.0/removebg", utils.Env().REMOVE_BG_API_URL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		utils.MessageWithReply(s, m, "Error creating request", logger)
		logger.Error().Err(err).Msg("Error creating request")
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", utils.Env().REMOVE_BG_API_KEY)

	response, err := client.Do(req)
	if err != nil {
		utils.MessageWithReply(s, m, "Error sending request to API", logger)
		logger.Error().Err(err).Msg("Error sending request to API")
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		utils.MessageWithReply(s, m, fmt.Sprintf("API Error: %s", response.Status), logger)
		logger.Error().Int("status", response.StatusCode).Str("body", string(body)).Msg("API Error")
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		utils.MessageWithReply(s, m, "Error reading response body", logger)
		logger.Error().Err(err).Msg("Error reading response body")
		return
	}

	var apiResponse entities.RemoveBgResponse
	var apiResponseAlt entities.RemoveBgResponseAlt
	var resultBase64 string

	if err := json.Unmarshal(body, &apiResponse); err == nil && apiResponse.Data != nil && apiResponse.Data.ResultB64 != "" {
		resultBase64 = apiResponse.Data.ResultB64
	} else {
		if err := json.Unmarshal(body, &apiResponseAlt); err == nil && apiResponseAlt.Result != "" {
			resultBase64 = apiResponseAlt.Result
		} else {
			utils.MessageWithReply(s, m, "Error parsing API response", logger)
			logger.Error().Err(err).Msg("Error parsing API response")
			return
		}
	}

	processedImageBytes, err := base64.StdEncoding.DecodeString(resultBase64)
	if err != nil {
		utils.MessageWithReply(s, m, "Error decoding processed image", logger)
		logger.Error().Err(err).Msg("Error decoding processed image")
		return
	}

	_, err = s.ChannelFileSendWithMessage(
		m.ChannelID,
		fmt.Sprintf("Background diubah menjadi warna: **%s**", command),
		"processed_image.png",
		bytes.NewReader(processedImageBytes),
	)
	if err != nil {
		utils.MessageWithReply(s, m, "Error sending processed image", logger)
		logger.Error().Err(err).Msg("Error sending processed image")
		return
	}
}
