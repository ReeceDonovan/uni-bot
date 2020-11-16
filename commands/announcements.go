package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ReeceDonovan/uni-bot/api"
	"github.com/UCCNetsoc/discord-bot/embed"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	rawCourses, err := api.QueryCourse(token)
	if err != nil {
		fmt.Println("Token error")
		s.ChannelMessageSend(channel, "Error querying Canvas API, is your token correct?")
		return
	}

	emb := embed.NewEmbed()
	emb.SetColor(0xab0df9)
	p := message.NewPrinter(language.English)
	body := ""
	emb.SetTitle("Assignments")
	for _, course := range rawCourses {
		assData := api.QueryAssign(strconv.Itoa(course.ID), token)
		if len(assData) >= 1 {
			body += p.Sprintf("__**%s**__\n\n", course.Name)
			for _, ass := range assData {
				body += p.Sprintf("%s ", ass.Name)
				body += p.Sprintf("[%s]\n", (ass.DueAt.UTC().Format("15:04 - 02/01")))
				body += p.Sprintf("Due in: **%d Days, ", int((time.Until(ass.DueAt)).Hours()/24))
				body += p.Sprintf("%d Hours**\n", int((((time.Until(ass.DueAt)).Hours()/24)-(float64(int((time.Until(ass.DueAt)).Hours()/24))))*24))
				body += p.Sprintf("%s\n\n", ass.HTMLURL)
			}
		}
	}
	emb.SetDescription(body)
	s.ChannelMessageSendEmbed(channel, emb.MessageEmbed)
}
