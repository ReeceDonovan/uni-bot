package commands

import (
	"context"
	"fmt"

	"github.com/Strum355/log"
	"github.com/UCCNetsoc/discord-bot/embed"
	"github.com/bwmarrin/discordgo"
)

func helpCommand(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	emb := embed.NewEmbed()
	emb.SetTitle("Uni-Bot Commands")
	description := ""

	for k, v := range helpMsgs {
		description += fmt.Sprintf("**`!%s`**: %s\n\n", k, v)
	}

	emb.SetDescription(description)
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to send help message")
		return
	}
}
