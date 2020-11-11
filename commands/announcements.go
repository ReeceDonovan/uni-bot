package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ReeceDonovan/CS-bot/api"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
)

func RefeshSchedule(schedule *gocron.Scheduler, s *discordgo.Session) {
	schedule.Clear()
	//TODO: Scheduler
	// AnnounceAssignments(s)
	// api.QueryCourse()

	// upcomingAssignment := api.QueryCanvas()

	// if len(upcomingAssignment) > 0 {
	// 	schedule.Every(1).Hour().StartAt((time.Unix((upcomingAssignment[0].Date - 600), 0))).Do(UpcomingAssignmentAnnounce, context.TODO(), s)
	// }
}

func AnnounceMsgHandler(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	token := viper.GetString("canvas.token")
	fmt.Println(m.GuildID)
	if (m.GuildID != strconv.Itoa(624287001335562289)) || (m.GuildID != strconv.Itoa(471231303317192733)) {
		commandStr, body := extractCommand(m.Content)
		fmt.Println(commandStr)
		if len(strings.Fields(body)) > 1 {
			token = strings.Fields(body)[1]
		}
	}
	AnnounceAssignments(s, m.ChannelID, token)
}

func AnnounceAssignments(s *discordgo.Session, channel string, token string) {
	message := ""
	rawCourses, err := api.QueryCourse(token)
	if err != nil {
		fmt.Println("Token error")
		s.ChannelMessageSend(channel, "Error querying Canvas API, is your token correct?")
		return
	}

	for _, course := range rawCourses {
		assData := api.QueryAssign(strconv.Itoa(course.ID), token)
		if len(assData) >= 1 {
			message += "**__" + course.Name + "__**\n\n"
			for _, ass := range assData {
				message += "**" + ass.Name + "**" + "\n"
				message += "Due at: " + ass.DueAt.UTC().Format("3:04 PM - Jan 2") + "\n"
				message += ass.HTMLURL
				message += "\n\n"
			}
		}
	}
	s.ChannelMessageSend(channel, message)
}
