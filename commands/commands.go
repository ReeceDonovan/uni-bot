package commands

import (
	"fmt"
	"log"
	"strings"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/ReeceDonovan/uni-bot/request"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// TODO: Add functions for more bot commands e.g Stats for specific modules

// Test command for now, sends basic string of modules/assignments from current term as message. Will eventually be an embed and not hardcoded termID

func TermAssignments(s *discordgo.Session, m *discordgo.MessageCreate) {

	CourseAssignment := request.QueryAssignments()

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	p := message.NewPrinter(language.English)

	body := ""

	emb.SetTitle("This Terms Assignments")

	for _, course := range CourseAssignment.Data.AllCourses {
		log.Println(course.CourseName)

		if course.Term.ID != "44" || course.EnrollmentsConnection == nil {
			continue
		}
		log.Println(course.CourseName)

		body += p.Sprintf("__**%s**__\n\n", course.CourseName)
		for _, assignment := range course.AssignmentsConnection.Nodes {
			body += p.Sprintf("%s \n", assignment.Name)
			body += p.Sprintf("%d ", int(assignment.ScoreStatistics.Max))
			body += p.Sprintf("%d ", int(assignment.ScoreStatistics.Mean))
			body += p.Sprintf("%d \n", int(assignment.ScoreStatistics.Min))
			// body += p.Sprintf("[%s]\n", (assignment.DueAt.UTC().Format("15:04 - 02/01")))
			// body += p.Sprintf("\n%s\n\n", assignment.HTMLURL)
		}
	}

	emb.SetDescription(body)
	s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)

}

func CourseStats(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, slug := extractCommand(m.Content)
	log.Println(slug)
	slug = strings.ToUpper(strings.Split(slug, " ")[1])
	CourseAssignment := request.QueryAssignments()

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle(slug + " Grade Statistics")
	for _, course := range CourseAssignment.Data.AllCourses {
		if course.Term.ID != "44" || course.EnrollmentsConnection == nil || course.CourseCode[len(course.CourseCode)-6:] != slug {
			continue
		}
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.ScoreStatistics.Max == 0 {
				emb.AddField("Max:", "-")
				emb.AddField("Mean:", "-")
				emb.AddField("Min:", "-")
			} else {
				emb.AddField("Max:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Max)))
				emb.AddField("Mean:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Mean)))
				emb.AddField("Min:", fmt.Sprintf("%d", int(assignment.ScoreStatistics.Min)))
			}
		}
	}
	emb.InlineAllFields()
	s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}
