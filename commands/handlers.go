package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/ReeceDonovan/uni-bot/middleware"
	"github.com/ReeceDonovan/uni-bot/models"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"link": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			discordUser, err := middleware.ValidateScope(s, i)
			if err != nil {
				ErrorHandler(s, i, err)
				return
			}
			scope := i.ApplicationCommandData().Options[0].StringValue()
			token := i.ApplicationCommandData().Options[1].StringValue()

			if len(token) < 8 {
				ErrorHandler(s, i, errors.New("invalid token, please try again"))
				return
			}

			if scope == "server" {

				server := &models.Server{
					SID: i.GuildID,
				}

				err := server.Get()
				if err == nil {
					ErrorHandler(s, i, errors.New("this server has already been linked, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					log.Printf("Error creating server link: %v", err)
					ErrorHandler(s, i, errors.New("error creating server link, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				}

				server = &models.Server{
					SID:         i.GuildID,
					CanvasToken: token,
				}

				err = server.Create()
				if err != nil {
					log.Printf("Error creating server link: %v", err)
					ErrorHandler(s, i, errors.New("error creating server link, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				}

			} else {

				user := &models.User{
					UID: discordUser.ID,
				}

				err := user.Get()
				if err == nil {
					ErrorHandler(s, i, errors.New("this user has already been linked, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					log.Printf("Error creating user link: %v", err)
					ErrorHandler(s, i, errors.New("error creating user link, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				}

				user = &models.User{
					UID:         discordUser.ID,
					CanvasToken: token,
				}

				err = user.Create()
				if err != nil {
					log.Printf("Error creating user link: %v", err)
					ErrorHandler(s, i, errors.New("error creating user link, please contact: <@342150581554774018> | Nõ̷̋t̴̏͆ĥ̵̆i̴̓̌c̵͌̎#9999"))
					return
				}
			}

			msgformat :=
				` Token has been linked:
				> Scope: %s
				> Token: ||%s||
`
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: 1 << 6,
					Content: fmt.Sprintf(
						msgformat,
						scope, token,
					),
				},
			})
			if scope == "server" {
				s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
					Content: fmt.Sprintf(
						`<@%s> has just linked this server, link your personal token using the /link command for personalised canvas data`, discordUser.ID,
					),
				})

			}

		},

		"assignments": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			discordUser, err := middleware.ValidateScope(s, i)
			if err != nil {
				ErrorHandler(s, i, err)
				return
			}
			scope := i.ApplicationCommandData().Options[0].StringValue()

			if scope == "server" {

				server := &models.Server{
					SID: i.GuildID,
				}

				err := server.Get()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Printf("Error getting server data: %v", err)
					ErrorHandler(s, i, errors.New("error getting server data, make sure the server has already been linked using the /link command"))
				}

				// TODO: Canvas API

			} else {

				user := &models.User{
					UID: discordUser.ID,
				}

				err := user.Get()
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Printf("Error getting user data: %v", err)
					ErrorHandler(s, i, errors.New("error getting user data, make sure the you have already linked a canvas token using the /link command"))
				}

			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						` Reply with assignment embed using either the user's stored token or the server wide token:
						> Scope: %s`,
						scope,
					),
				},
			})
		},
	}
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
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
