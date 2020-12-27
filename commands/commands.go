package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/ReeceDonovan/uni-bot/request"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// TODO: Add functions for more bot commands e.g Stats for specific modules

// Test command for now, sends basic string of modules/assignments from current term as message. Will eventually be an embed and not hardcoded termID

func CurrentAssignments(s *discordgo.Session, m *discordgo.MessageCreate) {
	CourseAssignment := request.QueryAssignments()
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Active Assignments")
	p := message.NewPrinter(language.English)
	body := ""
	for _, course := range CourseAssignment.Data.AllCourses {
		if course.Term.ID != "44" || course.EnrollmentsConnection == nil {
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
		if course.Term.ID != "44" || course.EnrollmentsConnection == nil || course.CourseCode[len(course.CourseCode)-6:] != slug {
			continue
		}
		for x, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.ScoreStatistics.Max == 0 {
				continue
			}
			valid = true
			emb.AddField("⠀", "__**CA:**__")
			emb.AddField("⠀", "**"+fmt.Sprintf("%d", x+1)+"**")
			emb.AddField("⠀", "⠀")
			emb.AddField("Max:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Max)))
			emb.AddField("Mean:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Mean)))
			emb.AddField("Min:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Min)))
		}
	}
	emb.InlineAllFields()
	if valid {
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Error getting module data")
	}
}
