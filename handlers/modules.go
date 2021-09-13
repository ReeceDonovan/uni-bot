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

func modules(s *discordgo.Session, i *discordgo.InteractionCreate) {
	discordUser, err := middleware.ValidateScope(s, i)
	if err != nil {
		ErrorHandler(s, i, err)
		return
	}
	scope := i.ApplicationCommandData().Options[0].StringValue()

	var courses *models.Courses

	var response *discordgo.InteractionResponse

	switch scope {
	case "server":
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
				Embeds: []*discordgo.MessageEmbed{createModuleList(discordUser, courses)},
			},
		}
	case "user":
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
				Flags:  1 << 6,
				Embeds: []*discordgo.MessageEmbed{createModuleList(discordUser, courses)},
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

func createModuleList(discordUser *discordgo.User, courses *models.Courses) *discordgo.MessageEmbed {
	courseFields := []*discordgo.MessageEmbedField{}
	for _, course := range *courses {
		courseFields = append(courseFields, &discordgo.MessageEmbedField{
			Name:   course.Name,
			Value:  fmt.Sprintf(" | [Canvas](%s/courses/%d) | [UCC](https://www.ucc.ie/admin/registrar/modules/?mod=%s) | Enrolled: %d", viper.GetString("canvas.domain"), course.ID, course.CourseCode[5:], course.TotalStudents),
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: discordUser.Username,
		},
		Color:  0xab0df9, // Purple
		Fields: courseFields,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: discordUser.AvatarURL("2048"),
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     viper.GetString("canvas.term")[2:] + " Module List",
	}
	return embed
}
