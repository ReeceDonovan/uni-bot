package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/ReeceDonovan/uni-bot/request"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// TODO: Add more commands
// TODO: Cleanup command functions (maybe work out a helper func for created the message embed to avoid repeated code)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	emb := embed.NewEmbed()
	emb.SetColor(0xab0df9)
	emb.SetTitle("Uni-Bot Commands")

	for k, v := range helpMsgs {
		emb.AddField("!"+k, v)
	}
	s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}

func CurrentAssignments(s *discordgo.Session, m *discordgo.MessageCreate) {

	CourseAssignment := request.QueryAssignments()
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Active Assignments")
	p := message.NewPrinter(language.English)
	body := ""
	for _, course := range CourseAssignment.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil {
			continue
		}
		body += p.Sprintf("__**%s**__\n\n", course.CourseName)
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.DueAt.Unix() < time.Now().AddDate(0, 0, 0).Unix() {
				continue
			}
			days := int(time.Until(assignment.DueAt).Hours() / 24)
			hours := int(time.Until(assignment.DueAt).Hours() - float64(int(days*24)))
			minutes := int(time.Until(assignment.DueAt).Minutes() - float64(int(days*24*60)+int(hours*60)))
			body += p.Sprintf("%s ", assignment.Name)
			body += p.Sprintf("[%s]\n", (assignment.DueAt.UTC().Format("15:04 - 02/01")))
			body += p.Sprintf("Due in: **%d Days, ", days)
			body += p.Sprintf("%d Hours, ", hours)
			body += p.Sprintf("%d Minutes", minutes)
			body += p.Sprintf("**\n%s\n\n", assignment.HTMLURL)
		}
	}
	if valid {
		emb.SetDescription(body)
	} else {
		emb.SetDescription("__No Assignments Found__")
	}
	s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}

func CourseStats(s *discordgo.Session, m *discordgo.MessageCreate) {
	cm, slug := extractCommand(m.Content)
	log.Println(slug)
	if cm == slug {
		s.ChannelMessageSend(m.ChannelID, "Please enter a valid module code")
		return
	}
	slug = strings.ToUpper(strings.Split(slug, " ")[1])
	CourseAssignment := request.QueryAssignments()

	valid := false
	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Available Grade Statistics: " + slug)
	for _, course := range CourseAssignment.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil || course.CourseCode[len(course.CourseCode)-6:] != slug {
			continue
		}
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.ScoreStatistics.Max == 0 {
				continue
			}
			valid = true
			log.Println(assignment.Name)
			emb.AddField("Max:", fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Max)))
			emb.AddField("Mean:", fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Mean)))
			emb.AddField("Min:", fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Min)))
		}
	}
	emb.InlineAllFields()
	if valid {
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Error getting module data")
	}
}

func CoordinatorInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	CourseAssignment := request.QueryAssignments()
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Module Coordinator Info")

	p := message.NewPrinter(language.English)
	for _, course := range CourseAssignment.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil {
			continue
		}
		body := ""
		valid = true

		for _, enrolled := range course.EnrollmentsConnection.Nodes {
			if enrolled.Type == "TeacherEnrollment" {
				body += p.Sprintf("[%s]("+viper.GetString("canvas.domain")+"/courses/"+course.ID+"/users/"+enrolled.User.ID+")\n", enrolled.User.Name)
			}
		}
		emb.AddField(course.CourseName, body)
	}
	if valid {
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Error getting module data")
	}

}
