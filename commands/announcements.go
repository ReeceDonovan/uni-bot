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
	allAss := []api.ParsedAssignment{}
	token := viper.GetString("canvas.token")
	rawCourses, err := api.QueryCourse(token)
	if err != nil {
		fmt.Println("Token error")
		return
	}
	for _, course := range rawCourses {
		assData := api.QueryAssign(strconv.Itoa(course.ID), token)
		if len(assData) >= 1 {
			for _, ass := range assData {
				allAss = append(allAss, ass)
			}
		}
	}

	for _, aAss := range allAss {
		if (aAss.DueAt.Unix() - time.Now().Unix()) < 57600 {
			AnnounceDue(context.TODO(), s, aAss)
		}
	}
}

func AnnounceDue(ctx context.Context, s *discordgo.Session, a api.ParsedAssignment) {
	channelID := "781515702266888194"
	emb := embed.NewEmbed()
	emb.SetColor(0xab0df9)
	p := message.NewPrinter(language.English)
	body := ""
	emb.SetTitle("Due Today")
	days := int(time.Until(a.DueAt).Hours() / 24)
	hours := int(time.Until(a.DueAt).Hours() - float64(int(days*24)))
	minutes := int(time.Until(a.DueAt).Minutes() - float64(int(days*24*60)+int(hours*60)))
	body += p.Sprintf("%s ", a.Name)
	body += p.Sprintf("[%s]\n", (a.DueAt.UTC().Format("15:04 - 02/01")))
	body += p.Sprintf("Due in: **%d Days, ", days)
	body += p.Sprintf("%d Hours, ", hours)
	body += p.Sprintf("%d Minutes", minutes)
	body += p.Sprintf("**\n%s\n\n", a.HTMLURL)

	emb.SetDescription(body)
	s.ChannelMessageSend(channelID, "<@&631495001573949450>")
	s.ChannelMessageSendEmbed(channelID, emb.MessageEmbed)
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
				days := int(time.Until(ass.DueAt).Hours() / 24)
				hours := int(time.Until(ass.DueAt).Hours() - float64(int(days*24)))
				minutes := int(time.Until(ass.DueAt).Minutes() - float64(int(days*24*60)+int(hours*60)))
				body += p.Sprintf("%s ", ass.Name)
				body += p.Sprintf("[%s]\n", (ass.DueAt.UTC().Format("15:04 - 02/01")))
				body += p.Sprintf("Due in: **%d Days, ", days)
				body += p.Sprintf("%d Hours, ", hours)
				body += p.Sprintf("%d Minutes", minutes)
				body += p.Sprintf("**\n%s\n\n", ass.HTMLURL)
			}
		}
	}
	emb.SetDescription(body)
	s.ChannelMessageSendEmbed(channel, emb.MessageEmbed)
}
