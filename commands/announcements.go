package commands

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/reecedonovan/CS-bot/api"
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

func AnnounceHandler(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	AnnounceAssignments(s, "")
}

func AnnounceMsgHandler(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
	AnnounceAssignments(s, m.ChannelID)
}

func AnnounceAssignments(s *discordgo.Session, channel string) {
	message := ""
	for _, course := range api.QueryCourse() {
		assData := api.QueryAssign(strconv.Itoa(course.ID))
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
