package handlers

import (
	"errors"
	"fmt"
	"log"

	"github.com/ReeceDonovan/uni-bot/api"
	"github.com/ReeceDonovan/uni-bot/middleware"
	"github.com/ReeceDonovan/uni-bot/models"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func stats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	discordUser, err := middleware.ValidateScope(s, i)
	if err != nil {
		ErrorHandler(s, i, err)
		return
	}

	var response *discordgo.InteractionResponse

	switch i.ApplicationCommandData().Options[0].StringValue() {
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

		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    "Available stats:",
				Flags:      1 << 6,
				Components: *createModuleStatsList("u", user.CanvasToken),
			},
		}

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

		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    "Available stats:",
				Components: *createModuleStatsList("s", server.CanvasToken),
			},
		}

	}
	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
}

func statsComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse

	var discordUser *discordgo.User

	if i.User != nil {
		discordUser = i.User
	} else {
		discordUser = i.Member.User
	}

	log.Printf("Stats component used by: %v", discordUser.Username)

	data := i.MessageComponentData()

	switch data.Values[0] {
	case "null":

	default:
		moduleID := data.Values[0][2:]
		var token string
		if data.Values[0][0:1] == "u" {

			user := &models.User{
				UID: discordUser.ID,
			}

			err := user.Get()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("Error getting user data: %v", err)
				ErrorHandler(s, i, errors.New("error getting user data, make sure the you have already linked a canvas token using the /link command"))
				return
			}
			token = user.CanvasToken

			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Flags: 1 << 6,
					// Content: "**Module Stats**",
					Embeds: createModuleStatsEmbed(moduleID, token),
				},
			}
		} else {
			server := &models.Server{
				SID: i.GuildID,
			}

			err := server.Get()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("Error getting server data: %v", err)
				ErrorHandler(s, i, errors.New("error getting server data, make sure the server has already been linked using the /link command"))
				return
			}
			token = server.CanvasToken

			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					// Content: "**Module Stats**",
					Embeds: createModuleStatsEmbed(moduleID, token),
				},
			}
		}
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
}

func createModuleStatsList(scope string, token string) *[]discordgo.MessageComponent {
	courses := api.GetCourses(token)

	var moduleOptions []discordgo.SelectMenuOption
	moduleCount := 0

	for _, module := range *courses {
		if module.Enrollments[0].ComputedCurrentScore != nil {
			moduleCount++
			moduleOptions = append(moduleOptions, discordgo.SelectMenuOption{
				Label:       module.CourseCode[5:],
				Value:       fmt.Sprintf("%s.%d", scope, module.ID),
				Description: module.Name[12:],
			})
		}

	}
	if moduleCount == 0 {
		moduleOptions = append(moduleOptions, discordgo.SelectMenuOption{
			Label:       "No stats available",
			Value:       "null",
			Description: "No stats available",
		})
	}

	return &[]discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					CustomID:    "statsSelect",
					Placeholder: "Stats: Choose a module",
					Options:     moduleOptions,
				},
			},
		},
	}
}

func createModuleStatsEmbed(moduleID string, token string) []*discordgo.MessageEmbed {
	assignments := api.GetAssignments(moduleID, token)

	course := api.GetCourse(moduleID, token)

	totalPossible, totalMax, totalMean, totalMin := 0.0, 0.0, 0.0, 0.0

	var embs []*discordgo.MessageEmbed

	for _, assignment := range *assignments {

		var fields []*discordgo.MessageEmbedField
		if assignment.ScoreStatistics != nil {
			totalPossible += assignment.PointsPossible
			totalMax += assignment.ScoreStatistics.Max
			totalMean += assignment.ScoreStatistics.Mean
			totalMin += assignment.ScoreStatistics.Min
			fields = []*discordgo.MessageEmbedField{
				{
					Name:   "Highest Mark",
					Value:  fmt.Sprintf("%.2f", assignment.ScoreStatistics.Max),
					Inline: true,
				}, {
					Name:   "Average Mark",
					Value:  fmt.Sprintf("%.2f", assignment.ScoreStatistics.Mean),
					Inline: true,
				},
				{
					Name:   "Lowest Mark",
					Value:  fmt.Sprintf("%.2f", assignment.ScoreStatistics.Min),
					Inline: true,
				},
			}

			name := assignment.Name
			if len(name) > 38 {
				name = name[:35] + "..."

			}
			embs = append(embs, &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name: name,
					URL:  assignment.HTMLURL,
				},
				Color:       0xab0df9,
				Fields:      fields,
				Description: "ㅤ",
			})
		}
	}
	ret := []*discordgo.MessageEmbed{
		{
			Author: &discordgo.MessageEmbedAuthor{
				Name: course.Name[5:],
				URL:  fmt.Sprintf("%s/courses/%d", viper.GetString("canvas.domain"), course.ID),
			},
			Title: fmt.Sprintf("Total Grades (out of %.2f marks)", totalPossible),
			Color: 0xab0df9,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Highest Mark",
					Value:  fmt.Sprintf("%.2f", totalMax),
					Inline: true,
				}, {
					Name:   "Average Mark",
					Value:  fmt.Sprintf("%.2f", totalMean),
					Inline: true,
				}, {
					Name:   "Lowest Mark",
					Value:  fmt.Sprintf("%.2f", totalMin),
					Inline: true,
				},
			},
			Description: "ㅤ",
		},
	}
	ret = append(ret, embs...)
	return ret
}
