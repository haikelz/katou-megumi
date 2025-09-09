package utils

import (
	"katou-megumi/pkg/configs"
	entities "katou-megumi/pkg/entities/generated"
	"log"

	"github.com/spf13/viper"
)

func Env() entities.EnvVariables {
	v := configs.NewViper()

	_ = v.BindEnv("DATABASE_URL")
	_ = v.BindEnv("DISCORD_TOKEN")
	_ = v.BindEnv("REMOVE_BG_API_KEY")
	_ = v.BindEnv("REMOVE_BG_API_URL")
	_ = v.BindEnv("JOKES_API_URL")
	_ = v.BindEnv("ANIME_QUOTE_API_URL")
	_ = v.BindEnv("DISTRO_INFO_API_URL")
	_ = v.BindEnv("DOA_API_URL")
	_ = v.BindEnv("QURAN_API_URL")
	_ = v.BindEnv("IMAGE_API_URL")
	_ = v.BindEnv("GEMINI_API_KEY")
	_ = v.BindEnv("ASMAUL_HUSNA_API_URL")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config file: %v", err)
		}
	}

	return entities.EnvVariables{
		DISCORD_TOKEN:        v.GetString("DISCORD_TOKEN"),
		REMOVE_BG_API_KEY:    v.GetString("REMOVE_BG_API_KEY"),
		REMOVE_BG_API_URL:    v.GetString("REMOVE_BG_API_URL"),
		JOKES_API_URL:        v.GetString("JOKES_API_URL"),
		ANIME_QUOTE_API_URL:  v.GetString("ANIME_QUOTE_API_URL"),
		DISTRO_INFO_API_URL:  v.GetString("DISTRO_INFO_API_URL"),
		DOA_API_URL:          v.GetString("DOA_API_URL"),
		QURAN_API_URL:        v.GetString("QURAN_API_URL"),
		IMAGE_API_URL:        v.GetString("IMAGE_API_URL"),
		GEMINI_API_KEY:       v.GetString("GEMINI_API_KEY"),
		ASMAUL_HUSNA_API_URL: v.GetString("ASMAUL_HUSNA_API_URL"),
	}
}
