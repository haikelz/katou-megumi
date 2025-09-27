package handlers

import (
	"encoding/json"
	"fmt"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func DoaHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	wg := &sync.WaitGroup{}
	ch := make(chan utils.HttpGetResponse, 1)
	go utils.Get(utils.Env().DOA_API_URL+"/api/doa/v1/random", s, m, logger, wg, ch)
	wg.Wait()
	close(ch)

	response := <-ch
	if response.Error != nil {
		utils.MessageWithReply(s, m, "Error fetching do'a", logger)
		logger.Error("Error fetching do'a", zap.Error(response.Error))
		return
	}

	var doaResponse []entities.Doa
	err := json.Unmarshal(response.Body, &doaResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Error unmarshalling do'a", logger)
		logger.Error("Error unmarshalling do'a", zap.Error(err))
		return
	}

	utils.MessageWithReply(s, m, fmt.Sprintf("%d - %s - %s - %s", doaResponse[0].Id, doaResponse[0].Doa, doaResponse[0].Ayat, doaResponse[0].Artinya, command), logger)
}
