package middleware

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func ValidateScope(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.User, error) {
	if i.ApplicationCommandData().Options[0].StringValue() == "server" {
		if len(i.GuildID) < 5 {
			return nil, errors.New("command used from an invalid server")
		}
	}
	var discordUser *discordgo.User

	if i.User != nil {
		discordUser = i.User
	} else if i.Member != nil {
		discordUser = i.Member.User
	} else {
		return nil, errors.New("error creating user link: User data nil")
	}

	return discordUser, nil
}
