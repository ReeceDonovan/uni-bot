package commands

import (
	"log"
	"strings"

	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/bwmarrin/discordgo"
)

func Link(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, slug := extractCommand(m.Content)
	log.Println(strings.Fields(slug)[1])
	info := strings.Fields(slug)[1:]
	config.UpdateData(&config.ServerData{info[0], info[1], info[2]})
}
