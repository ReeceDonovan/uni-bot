package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"link": link,

		// TODO: Assignments
		"assignments": assignments,

		"modules": modules,
		"stats":   stats,
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

}

func ErrorHandler(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	log.Println(err.Error())
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: err.Error(),
		},
	})
}
