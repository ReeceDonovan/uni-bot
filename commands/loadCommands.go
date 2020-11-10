package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	commandsMap = make(map[string]func(context.Context, *discordgo.Session, *discordgo.MessageCreate))
)

type commandFunc func(context.Context, *discordgo.Session, *discordgo.MessageCreate)

func command(name string, helpMessage string, function commandFunc) {
	commandsMap[name] = function
}

func Register(s *discordgo.Session) {
	command(
		"assignment",
		"replies with embed of the the next upcoming netsoc event, queried from facebook",
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
		fmt.Println(viper.GetString("bot.prefix"))
		fmt.Println("command check")
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
