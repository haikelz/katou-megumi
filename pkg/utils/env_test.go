package utils

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	// Global config viper are not works in testing environment. Provide your .env first, place in this folder!
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	err := v.ReadInConfig()
	assert.NoError(t, err)

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

	env := Env()

	// check every env variables
	assert.Contains(t, env.ASMAUL_HUSNA_API_URL, v.GetString("ASMAUL_HUSNA_API_URL"))
	assert.Contains(t, env.GEMINI_API_KEY, v.GetString("GEMINI_API_KEY"))
	assert.Contains(t, env.IMAGE_API_URL, v.GetString("IMAGE_API_URL"))
	assert.Contains(t, env.QURAN_API_URL, v.GetString("QURAN_API_URL"))
	assert.Contains(t, env.DOA_API_URL, v.GetString("DOA_API_URL"))
	assert.Contains(t, env.DISTRO_INFO_API_URL, v.GetString("DISTRO_INFO_API_URL"))
	assert.Contains(t, env.ANIME_QUOTE_API_URL, v.GetString("ANIME_QUOTE_API_URL"))
	assert.Contains(t, env.JOKES_API_URL, v.GetString("JOKES_API_URL"))
	assert.Contains(t, env.REMOVE_BG_API_KEY, v.GetString("REMOVE_BG_API_KEY"))
	assert.Contains(t, env.REMOVE_BG_API_URL, v.GetString("REMOVE_BG_API_URL"))
	assert.Contains(t, env.DISCORD_TOKEN, v.GetString("DISCORD_TOKEN"))
}
