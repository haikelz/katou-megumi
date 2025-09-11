package main

import (
	"katou-megumi/pkg/configs"
	"katou-megumi/pkg/handlers"
	"katou-megumi/pkg/utils"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	logger := configs.NewZap()

	discordToken := utils.Env().DISCORD_TOKEN
	discord := configs.NewDiscord(discordToken)

	discord.Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", r.User.Username, r.User.Discriminator)
	})

	discord.Client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		var splitMessage []string = strings.Split(m.Content, " ")
		var args []string = splitMessage[1:]
		var command string = strings.Join(args, " ")

		if splitMessage[0] == "!info" {
			s.ChannelTyping(m.ChannelID)
			handlers.InfoHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ping" {
			s.ChannelTyping(m.ChannelID)
			handlers.PingHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!salam" {
			s.ChannelTyping(m.ChannelID)
			handlers.SalamHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!asmaulhusna" {
			s.ChannelTyping(m.ChannelID)
			handlers.AsmaulHusnaHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ask" {
			s.ChannelTyping(m.ChannelID)
			handlers.GeminiHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!jokes" {
			s.ChannelTyping(m.ChannelID)
			handlers.JokeHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!jadwalsholat" {
			s.ChannelTyping(m.ChannelID)
			handlers.JadwalSholatHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!doa" {
			s.ChannelTyping(m.ChannelID)
			handlers.DoaHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!quote" {
			s.ChannelTyping(m.ChannelID)
			handlers.QuoteHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!editbackground" {
			s.ChannelTyping(m.ChannelID)
			handlers.BackgroundPhotoHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!ocr" {
			s.ChannelTyping(m.ChannelID)
			handlers.OcrHandler(s, m, logger, command)
		}
		if splitMessage[0] == "!shutdown" {
			s.ChannelTyping(m.ChannelID)
			handlers.ShutdownHandler(s, m, logger, command)
		}
	})

	discord.Client.Open()

	defer discord.Client.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
