package commands

import (
	"context"
	"fmt"
	"strconv"

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

func AnnounceHandler(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, token string) {
	AnnounceAssignments(s, "", token)
}

func AnnounceMsgHandler(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	var token string
	if len(m.GuildID) == 0 {
		// ctx := context.WithValue(context.Background(), log.Key, log.Fields{
		// 	"author_id":  m.Author.ID,
		// 	"channel_id": m.ChannelID,
		// 	"guild_id":   "DM",
		// })
		fmt.Println(m.Content)
		token = m.Content
	} else {
		token = viper.GetString("canvas.token")
	}
	AnnounceAssignments(s, m.ChannelID, token)
}

func AnnounceAssignments(s *discordgo.Session, channel string, token string) {
	message := ""
	for _, course := range api.QueryCourse(token) {
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
