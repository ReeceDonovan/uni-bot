package commands

import (
	"context"
	"strings"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	helpMsgs    = make(map[string]string)
	commandsMap = make(map[string]func(context.Context, *discordgo.Session, *discordgo.MessageCreate))
)

type commandFunc func(context.Context, *discordgo.Session, *discordgo.MessageCreate)

func command(name string, helpMessage string, function commandFunc) {
	helpMsgs[name] = helpMessage
	commandsMap[name] = function
}

func Register(s *discordgo.Session) {
	command("help", "Replies list of commands", helpCommand)
	command(
		"assignment",
		"Replies with formatted embed of active course assignments, parsed from Canvas API",
		AnnounceMsgHandler,
	)
	s.AddHandler(messageCreate)
}

// Called whenever a message is sent in a server the bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if !strings.HasPrefix(m.Content, viper.GetString("bot.prefix")) {
		return
	}
	callCommand(s, m)
}

func callCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	ctx := context.Background()
	commandStr, body := extractCommand(m.Content)
	// if command is a normal command
	if command, ok := commandsMap[commandStr]; ok {
		ctx := context.WithValue(ctx, log.Key, log.Fields{
			"author_id":  m.Author.ID,
			"channel_id": m.ChannelID,
			"guild_id":   m.GuildID,
			"command":    commandStr,
			"body":       body,
		})
		log.WithContext(ctx).Info("invoking standard command")
		command(ctx, s, m)
		return
	}
}

func extractCommand(c string) (commandStr string, body string) {
	body = strings.TrimPrefix(c, viper.GetString("bot.prefix"))
	commandStr = strings.Fields(body)[0]
	return
}
