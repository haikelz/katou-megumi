package handlers

import (
	"katou-megumi/pkg/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func InfoHandler(s *discordgo.Session, m *discordgo.MessageCreate, logger *zerolog.Logger, command string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var builder strings.Builder

	builder.WriteString(`
# 🤖 Katou Megumi Bot
*My Istri Discord Bot*

---

## 📋 **Daftar Perintah**

- !salam - Memberikan salam kepada pengguna
- !info - Menampilkan informasi tentang bot
- !ask - Mengajukan pertanyaan kepada bot
- !editbackground - Mengedit background profil
- !sticker - Mengirim sticker
- !jadwalsholat - Menampilkan jadwal sholat
- !doa - Menampilkan doa-doa harian
- !asmaulhusna - Menampilkan Asmaul Husna
- !jokes - Menampilkan lelucon
- !animequote - Quote anime random
- !ocr - Membaca teks dari gambar (Optical Character Recognition)
- !shutdown - Menutup bot

---

## 👨‍💻 **Developer**

[haikelz](https://github.com/haikelz/)

---

*Selamat menggunakan Katou Megumi Bot! 💜*	
`)

	katouMegumiImage := "https://avatars.githubusercontent.com/u/77146709?v=4"

	var katouMegumiImageEmbed = &discordgo.MessageEmbed{
		Description: builder.String(),
		Image:       &discordgo.MessageEmbedImage{URL: katouMegumiImage},
	}
	utils.MessageWithEmbedReply(s, m, katouMegumiImageEmbed, logger)
}
