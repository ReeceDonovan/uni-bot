package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var commandIDs map[string]string
var (
	commands = []discordgo.ApplicationCommand{
		{
			Name:        "link",
			Description: "Command for linking a canvas token",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "scope",
					Description: "Scope",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "User",
							Value: "user",
						},
						{
							Name:  "Server",
							Value: "server",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "token",
					Description: "Canvas Token",
					Required:    true,
				},
			},
		},
		{
			Name:        "assignments",
			Description: "Command for getting upcoming/active assignments",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "scope",
					Description: "Scope",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "User",
							Value: "user",
						},
						{
							Name:  "Server",
							Value: "server",
						},
					},
				},
			},
		},
		{
			Name: "stats",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "scope",
					Description: "Scope",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "User",
							Value: "user",
						},
						{
							Name:  "Server",
							Value: "server",
						},
					},
				},
			},
			Description: "Get available grades stats for modules",
		},
		{
			Name: "modules",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "scope",
					Description: "Scope",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "User",
							Value: "user",
						},
						{
							Name:  "Server",
							Value: "server",
						},
					},
				},
			},
			Description: "Get list of modules",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"link": link,

		// TODO: Assignments
		"assignments": assignments,

		"modules": modules,

		"stats": stats,
	}
	componentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"statsSelect": statsComponent,
	}
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	commandIDs = make(map[string]string, len(commands))

	for _, cmd := range commands {
		rcmd, err := s.ApplicationCommandCreate(viper.GetString("discord.app"), viper.GetString("discord.guild"), &cmd)
		if err != nil {
			log.Fatalf("Cannot create slash command %q: %v", cmd.Name, err)
		}

		commandIDs[rcmd.ID] = rcmd.Name
	}
}

func ErrorHandler(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: err.Error(),
		},
	})
}
