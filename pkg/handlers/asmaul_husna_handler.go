package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func AsmaulHusnaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if command == "" {
		asmaulHusnaResponse := getAllAsmaulHusna(utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse == nil {
			utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
			logger.Error("Error fetching Asmaul Husna", zap.Error(errors.New("error fetching Asmaul Husna")))
			return
		}

		loopAsmaulHusnaMessage(0, 20, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(21, 40, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(41, 60, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(61, 80, asmaulHusnaResponse, s, m, logger)
		loopAsmaulHusnaMessage(81, 98, asmaulHusnaResponse, s, m, logger)

		return
	}

	if number, err := strconv.Atoi(command); err == nil {
		asmaulHusnaResponse := getAsmaulHusnaByUrutan(number, utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse == nil || asmaulHusnaResponse.Data.Urutan == 0 {
			utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
			logger.Error("Asmaul Husna tidak ditemukan", zap.Error(errors.New("Asmaul Husna tidak ditemukan"+command)))
			return
		}

		utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
		return
	}

	asmaulHusnaResponse := getAsmaulHusnaByLatin(command, utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

	if asmaulHusnaResponse == nil || asmaulHusnaResponse.Data.Urutan == 0 {
		utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
		logger.Error("Asmaul Husna tidak ditemukan", zap.Error(errors.New("Asmaul Husna tidak ditemukan"+command)))
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
}

func loopAsmaulHusnaMessage(start int, end int, asmaulHusnaResponse []*entities.AsmaulHusna, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) {
	content := ""

	for _, v := range asmaulHusnaResponse[start:end] {
		content += fmt.Sprintf("%d - %s - %s - %s\n", v.Urutan, v.Latin, v.Arab, v.Arti)
	}

	utils.MessageWithReply(s, m, content, logger)
}

func getAllAsmaulHusna(ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) []*entities.AsmaulHusna {
	body := utils.Get(ASMAUL_HUSNA_API_URL+"/api/all", s, m, logger)

	var asmaulHusnaResponse entities.AsmaulHusnaResponse
	err := json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		logger.Error("Error unmarshalling Asmaul Husna", zap.Error(err))
		return nil
	}

	return asmaulHusnaResponse.Data
}

func getAsmaulHusnaByUrutan(urutan int, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) *entities.AsmaulHusnaByLatinOrUrutanResponse {
	body := utils.Get(ASMAUL_HUSNA_API_URL+"/api/"+strconv.Itoa(urutan), s, m, logger)

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err := json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asma'ul Husna", logger)
		logger.Error("Error unmarshalling Asmaul Husna", zap.Error(err))
		return nil
	}

	return &asmaulHusnaResponse
}

func getAsmaulHusnaByLatin(latin string, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger) *entities.AsmaulHusnaByLatinOrUrutanResponse {
	body := utils.Get(ASMAUL_HUSNA_API_URL+"/api/latin/"+latin, s, m, logger)

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err := json.Unmarshal(body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		return nil
	}

	return &asmaulHusnaResponse
}
