package handlers

import (
	"errors"
	"fmt"
	"log"

	"sort"
	"sync"

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
		log.Printf("Stats error: %v", err)
		log.Printf("Response: %v", *response)
		log.Printf("Interaction: %v", *i.Interaction)
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
		log.Printf("Stats component error: %v", err)
		log.Printf("Response: %v", response)
		log.Printf("Interaction: %v", i)
	}
}

func createModuleStatsList(scope string, token string) *[]discordgo.MessageComponent {
	courses := *api.GetCourses(token)

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})

	if len(courses) > 30 {
		courses = courses[len(courses)-30:]
	}

	var validCourses []models.Course

	var wg sync.WaitGroup
	for _, course := range courses {
		wg.Add(1)
		go func(course models.Course) {
			defer wg.Done()
			assignments := api.GetAssignments(fmt.Sprintf("%v", course.ID), token)
			for _, assignment := range *assignments {
				if assignment.ScoreStatistics != nil && assignment.ScoreStatistics.Mean > 0 {
					validCourses = append(validCourses, course)
					break
				}
			}

		}(course)
	}
	wg.Wait()

	var moduleOptions []discordgo.SelectMenuOption
	for _, course := range validCourses {
		moduleOptions = append(moduleOptions, discordgo.SelectMenuOption{
			Label:       course.CourseCode[5:],
			Value:       fmt.Sprintf("%s.%d", scope, course.ID),
			Description: course.Name[12:],
		})
	}

	sort.Slice(moduleOptions, func(i, j int) bool {
		return moduleOptions[i].Label < moduleOptions[j].Label
	})

	if len(moduleOptions) > 25 {
		moduleOptions = moduleOptions[len(moduleOptions)-25:]
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
			if len(name) > 35 {
				name = name[:32] + "..."
			}
			name = fmt.Sprintf("%s (%.2f)", name, assignment.PointsPossible)
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
