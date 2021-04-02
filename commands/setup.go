package commands

import (
	"log"
	"strings"

	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type commandFunc func(session *discordgo.Session, message *discordgo.MessageCreate)

var commandsMap = make(map[string]func(*discordgo.Session, *discordgo.MessageCreate))
var helpMsgs = make(map[string]string)

func command(name string, helpMessage string, function commandFunc) {
	helpMsgs[name] = helpMessage
	commandsMap[name] = function
}

func RegisterCommands(s *discordgo.Session) {
	command("help", "List available Uni-Bot commands", HelpCommand)
	command("assignment", "List active course assignments", Assignments)
	command("stats", "List grade statistics from specified module e.g !stats CS2502", CourseStats)
	command("contact", "List canvas user page for Professors and Course Coordinators", CoordinatorInfo)
	command("modules", "List canvas modules page for the current year", ModuleList)
	command("link", "[Owner Only] Link this server and channel with canvas token e.g !link <Token>", Link)
	s.AddHandler(messageCreate)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if !strings.HasPrefix(m.Content, viper.GetString("discord.prefix")) {
		return
	}
	callCommand(s, m)
}

func callCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID == "" {
		s.ChannelMessageSend(m.ChannelID, "> **Please use this bot in a properly linked server**")
		return
	}
	commandStr, _ := extractCommand(m.Content)
	if command, ok := commandsMap[commandStr]; ok {
		log.Println("Command Triggered")
		sr := viper.Get("servers.active").([]config.ServerData)
		for _, ser := range sr {
			if m.GuildID == ser.ServerID || commandStr == "link" {
				command(s, m)
				return
			}
		}
		s.ChannelMessageSend(m.ChannelID, "> **Server is not linked. Have the server owner use the !link command with a canvas account token in the channel you wish to receive assignment alerts**")
	}
}

func extractCommand(c string) (commandStr string, body string) {
	body = strings.TrimPrefix(c, viper.GetString("discord.prefix"))
	commandStr = strings.Fields(body)[0]
	return
}
