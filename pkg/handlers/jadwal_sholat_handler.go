package handlers

import (
	"encoding/json"
	entities "katou-megumi/pkg/entities/generated"
	"katou-megumi/pkg/utils"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func JadwalSholatHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zap.Logger, command string) {
	if len(command) <= 2 {
		if len(command) == 0 {
			logger.Info("User" + m.Author.Username + "is requesting jadwal sholat by default")
			utils.MessageWithReply(s, m, "Ini adalah perintah untuk mendapatkan jadwal sholat sesuai dengan nama daerah yang dimasukkan. Cukup ketik *!jadwalsholat <your_daerah>*", logger)
		} else {
			logger.Error("Get jadwalsholat failed " + command + " from " + m.Author.Username)
			utils.MessageWithReply(s, m, "Maaf, panjang karakter daerah yang dimasukkan tidak boleh kurang dari atau sama dengan 2!", logger)
		}
	}

	cityId, err := getCityId(s, m, command, logger)
	if err != nil {
		logger.Error("ID Kota tidak ditemukan!", zap.Error(err))
		utils.MessageWithReply(s, m, "Maaf, terjadi kesalahan saat mengambil data jadwal sholat!", logger)
		return
	}

	jadwalSholatResponse, err := getJadwalSholat(s, m, cityId, logger)
	if err != nil {
		utils.MessageWithReply(s, m, "Maaf, terjadi kesalahan saat mengambil data jadwal sholat!", logger)
		logger.Error("Data jadwal sholat tidak ditemukan!", zap.Error(err))
		return
	}

	today := time.Now().Format("2006-01-02")

	content := "Jadwal Sholat " + jadwalSholatResponse.Data.Daerah + " " + jadwalSholatResponse.Data.Lokasi + " " + today + "\n" +
		"Imsak: " + jadwalSholatResponse.Data.Jadwal.Imsak + "\n" +
		"Subuh: " + jadwalSholatResponse.Data.Jadwal.Subuh + "\n" +
		"Terbit: " + jadwalSholatResponse.Data.Jadwal.Terbit + "\n" +
		"Dhuha: " + jadwalSholatResponse.Data.Jadwal.Dhuha + "\n" +
		"Dzuhur: " + jadwalSholatResponse.Data.Jadwal.Dzuhur + "\n" +
		"Ashar: " + jadwalSholatResponse.Data.Jadwal.Ashar + "\n" +
		"Maghrib: " + jadwalSholatResponse.Data.Jadwal.Maghrib + "\n" +
		"Isya: " + jadwalSholatResponse.Data.Jadwal.Isya

	utils.MessageWithReply(s, m, content, logger)
}

func getCityId(s *discordgo.Session, m *discordgo.MessageCreate, cityName string, logger *zap.Logger) (string, error) {
	body := utils.Get(utils.Env().QURAN_API_URL+"/v2/sholat/kota/cari/"+cityName, s, m, logger)

	var jadwalSholatResponse entities.JadwalSholaCityIdResponse

	err := json.Unmarshal(body, &jadwalSholatResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Maaf, terjadi kesalahan saat mengambil data jadwal sholat!", logger)
		logger.Error("Error unmarshalling body", zap.Error(err))
		return "", err
	}

	return jadwalSholatResponse.Data[0].Id, nil
}

func getJadwalSholat(s *discordgo.Session, m *discordgo.MessageCreate, cityId string, logger *zap.Logger) (*entities.JadwalSholatResponse, error) {
	today := time.Now().Format(utils.TIME_FORMAT)

	body := utils.Get(utils.Env().QURAN_API_URL+"/v2/sholat/jadwal/"+cityId+"/"+today, s, m, logger)

	var jadwalSholatResponse entities.JadwalSholatResponse
	err := json.Unmarshal(body, &jadwalSholatResponse)
	if err != nil {
		utils.MessageWithReply(s, m, "Maaf, terjadi kesalahan saat mengambil data jadwal sholat!", logger)
		logger.Error("Error unmarshalling body", zap.Error(err))
		return nil, err
	}

	return &jadwalSholatResponse, nil
}
