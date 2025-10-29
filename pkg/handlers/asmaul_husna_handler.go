package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"strconv"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func AsmaulHusnaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	if command == "" {
		asmaulHusnaResponse := getAllAsmaulHusna(utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse == nil {
			utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
			logger.Error().Err(errors.New("error fetching Asmaul Husna")).Msg("Error fetching Asmaul Husna")
			return
		}

		l := Loop{}

		l.loopAsmaulHusnaMessage(0, 20, asmaulHusnaResponse, s, m, logger)
		l.loopAsmaulHusnaMessage(21, 40, asmaulHusnaResponse, s, m, logger)
		l.loopAsmaulHusnaMessage(41, 60, asmaulHusnaResponse, s, m, logger)
		l.loopAsmaulHusnaMessage(61, 80, asmaulHusnaResponse, s, m, logger)
		l.loopAsmaulHusnaMessage(81, 98, asmaulHusnaResponse, s, m, logger)

		return
	}

	if number, err := strconv.Atoi(command); err == nil {
		asmaulHusnaResponse := getAsmaulHusnaByUrutan(number, utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

		if asmaulHusnaResponse == nil || asmaulHusnaResponse.Data.Urutan == 0 {
			utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
			logger.Error().Err(errors.New("Asmaul Husna tidak ditemukan" + command)).Msg("Asmaul Husna tidak ditemukan")
			return
		}

		utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
		return
	}

	asmaulHusnaResponse := getAsmaulHusnaByLatin(command, utils.Env().ASMAUL_HUSNA_API_URL, s, m, logger)

	if asmaulHusnaResponse == nil || asmaulHusnaResponse.Data.Urutan == 0 {
		utils.MessageWithReply(s, m, "Asmaul Husna tidak ditemukan", logger)
		logger.Error().Err(errors.New("Asmaul Husna tidak ditemukan" + command)).Msg("Asmaul Husna tidak ditemukan")
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", asmaulHusnaResponse.Data.Urutan, asmaulHusnaResponse.Data.Latin, asmaulHusnaResponse.Data.Arab, asmaulHusnaResponse.Data.Arti), logger)
}

type Loop struct {
	mu sync.Mutex
}

func (l *Loop) loopAsmaulHusnaMessage(start int, end int, asmaulHusnaResponse []*entities.AsmaulHusna, s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) {
	var builder strings.Builder

	for _, v := range asmaulHusnaResponse[start:end] {
		builder.WriteString(fmt.Sprintf("%d - %s - %s - %s\n", v.Urutan, v.Latin, v.Arab, v.Arti))
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	utils.MessageWithReply(s, m, builder.String(), logger)
}

func getAllAsmaulHusna(ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) []*entities.AsmaulHusna {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)

	wg.Add(1)
	go utils.Get(ASMAUL_HUSNA_API_URL+"/api/all", s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		logger.Error().Err(response.Error).Msg("Error fetching Asmaul Husna")
		return nil
	}

	var asmaulHusnaResponse entities.AsmaulHusnaResponse
	err := json.Unmarshal(response.Body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		logger.Error().Err(err).Msg("Error unmarshalling Asmaul Husna")
		return nil
	}

	return asmaulHusnaResponse.Data
}

func getAsmaulHusnaByUrutan(urutan int, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) *entities.AsmaulHusnaByLatinOrUrutanResponse {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)

	go utils.Get(ASMAUL_HUSNA_API_URL+"/api/"+strconv.Itoa(urutan), s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		logger.Error().Err(response.Error).Msg("Error fetching Asmaul Husna")
		return nil
	}

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err := json.Unmarshal(response.Body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asma'ul Husna", logger)
		logger.Error().Err(err).Msg("Error unmarshalling Asmaul Husna")
		return nil
	}

	return &asmaulHusnaResponse
}

func getAsmaulHusnaByLatin(latin string, ASMAUL_HUSNA_API_URL string, s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger) *entities.AsmaulHusnaByLatinOrUrutanResponse {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)
	go utils.Get(ASMAUL_HUSNA_API_URL+"/api/latin/"+latin, s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		utils.MessageWithReply(s, m, "Error fetching Asmaul Husna", logger)
		logger.Error().Err(response.Error).Msg("Error fetching Asmaul Husna")
		return nil
	}

	var asmaulHusnaResponse entities.AsmaulHusnaByLatinOrUrutanResponse

	err := json.Unmarshal(response.Body, &asmaulHusnaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling Asmaul Husna", logger)
		return nil
	}

	return &asmaulHusnaResponse
}
