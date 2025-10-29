package handlers

import (
	"encoding/json"
	"fmt"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)
	go utils.Get(utils.Env().DOA_API_URL+"/api/doa/v1/random", s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		utils.MessageWithReply(s, m, "Error fetching do'a", logger)
		logger.Error().Err(response.Error).Msg("Error fetching do'a")
		return
	}

	var doaResponse []entities.Doa
	err := json.Unmarshal(response.Body, &doaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling do'a", logger)
		logger.Error().Err(err).Msg("Error unmarshalling do'a")
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", doaResponse[0].Id, doaResponse[0].Doa, doaResponse[0].Ayat, doaResponse[0].Artinya, command), logger)
}
