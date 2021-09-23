package handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ReeceDonovan/uni-bot/api"
	"github.com/ReeceDonovan/uni-bot/middleware"
	"github.com/ReeceDonovan/uni-bot/models"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func assignments(s *discordgo.Session, i *discordgo.InteractionCreate) {
	discordUser, err := middleware.ValidateScope(s, i)
	if err != nil {
		ErrorHandler(s, i, err)
		return
	}
	scope := i.ApplicationCommandData().Options[0].StringValue()

	var courses *models.Courses
	var response *discordgo.InteractionResponse

	if scope == "server" {

		server := &models.Server{
			SID: i.GuildID,
		}

		err := server.Get()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error getting server data: %v", err)
			ErrorHandler(s, i, errors.New("error getting server data, make sure the server has already been linked using the /link command"))
			return
		}

		courses = api.GetCourses(server.CanvasToken)

		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: createAssignmentList(discordUser, courses, server.CanvasToken),
			},
		}

	} else {

		user := &models.User{
			UID: discordUser.ID,
		}

		err := user.Get()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error getting user data: %v", err)
			ErrorHandler(s, i, errors.New("error getting user data, make sure the you have already linked a canvas token using the /link command"))
			return
		}

		courses = api.GetCourses(user.CanvasToken)

		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: createAssignmentList(discordUser, courses, user.CanvasToken),
				Flags:  1 << 6,
			},
		}

	}

	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Printf("Error responding data: %v", err)
		ErrorHandler(s, i, errors.New("error occurred"))
		return
	}
}

func createAssignmentList(discordUser *discordgo.User, courses *models.Courses, token string) (assignmentEmbeds []*discordgo.MessageEmbed) {
	for _, course := range *courses {
		if course.Term.Name != viper.GetString("canvas.term") {
			continue
		}
		assignments := api.GetAssignments(fmt.Sprintf("%d", course.ID), token)
		if assignments == nil || len(*assignments) < 1 {
			continue
		}
		var assignmentFields []*discordgo.MessageEmbedField
		valid := false
		for _, assignment := range *assignments {
			if assignment.DueAt.Unix() < time.Now().AddDate(0, 0, 0).Unix() {
				continue
			}
			valid = true
			marks := fmt.Sprintf("%.0f Marks | ", assignment.PointsPossible)
			countdown := fmt.Sprintf("<t:%d:R>\n[%s](%s)\n\n", assignment.DueAt.Unix(), assignment.DueAt.Format("02 Jan 2006 15:04"), assignment.HTMLURL)
			fields := &discordgo.MessageEmbedField{
				Name:   marks + assignment.Name,
				Value:  countdown,
				Inline: false,
			}
			assignmentFields = append(assignmentFields, fields)
		}
		if !valid {
			continue
		}
		embed := &discordgo.MessageEmbed{
			Color:     0xab0df9, // Purple
			Fields:    assignmentFields,
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     course.Name[5:],
		}
		assignmentEmbeds = append(assignmentEmbeds, embed)
	}
	if len(assignmentEmbeds) < 1 {
		assignmentEmbeds = append(assignmentEmbeds, &discordgo.MessageEmbed{
			Color:       0xab0df9,
			Description: "**No active assignments**",
			Timestamp:   time.Now().Format(time.RFC3339),
		})
	}
	return assignmentEmbeds
}
