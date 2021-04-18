package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/ReeceDonovan/uni-bot/api"
	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/bwmarrin/discordgo"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	emb := embed.NewEmbed()
	emb.SetColor(0xab0df9)
	emb.SetTitle("Uni-Bot Commands")

	for key, val := range helpMsgs {
		emb.AddField("!"+key, val)
	}
	s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}

func Link(s *discordgo.Session, m *discordgo.MessageCreate) {
	g, _ := s.State.Guild(m.GuildID)
	if g.OwnerID != m.Author.ID {
		s.ChannelMessageSend(m.ChannelID, "> **Only the server's owner can access this command**")
		return
	}
	_, slug := extractCommand(m.Content)
	if len(strings.Fields(slug)) != 2 {
		log.Println("Link command error: Insufficient details")
		s.ChannelMessageSend(m.ChannelID, "> **Link Failed: Please run command in channel you'd like to link for announcements. !link <CanvasToken>**")
		return
	}
	info := []string{m.GuildID, strings.Fields(slug)[1], m.ChannelID}
	upErr := config.UpdateData(&config.ServerData{ServerID: info[0], CanvasToken: info[1], AlertChannel: info[2]})
	if upErr != nil {
		log.Println(upErr)
		s.ChannelMessageSend(m.ChannelID, "> **Link Error: Server already linked**")
	} else {
		s.ChannelMessageSend(m.ChannelID, "> **Link Successful**")
	}
}

func Assignments(s *discordgo.Session, m *discordgo.MessageCreate) {

	courseAssignments := api.GetAssignments(m.GuildID)
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Active Assignments")
	p := message.NewPrinter(language.English)
	body := ""
	for _, course := range courseAssignments.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil {
			continue
		}
		assignmentsExist := false
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.DueAt.Unix() < time.Now().AddDate(0, 0, 0).Unix() {
				continue
			}
			if assignmentsExist == false {
				body += p.Sprintf("__**%s**__\n", course.CourseName[5:])
				valid, assignmentsExist = true, true
			}
			days := int(time.Until(assignment.DueAt).Hours() / 24)
			hours := int(time.Until(assignment.DueAt).Hours() - float64(int(days*24)))
			minutes := int(time.Until(assignment.DueAt).Minutes() - float64(int(days*24*60)+int(hours*60)))
			body += p.Sprintf("%.0f Marks\n", assignment.PointsPossible)
			body += p.Sprintf("[%s](%s)\n", assignment.Name, assignment.HTMLURL)
			body += p.Sprintf("**%d Days, ", days)
			body += p.Sprintf("%d Hours, ", hours)
			body += p.Sprintf("%d Minutes** 	|	", minutes)
			body += p.Sprintf("%s\n\n", (assignment.DueAt.Format("02 Jan 2006 15:04")))
		}
	}
	if valid {
		emb.SetDescription(body)
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "> **No Assignments Found**")
	}
}

func DueAssignments(s *discordgo.Session) {
	sr := viper.Get("servers.active").([]config.ServerData)
	for _, ser := range sr {
		dueAssignmentsHelper(s, ser.ServerID, ser.AlertChannel)
	}
}

func dueAssignmentsHelper(s *discordgo.Session, serverID string, alertChannelID string) {

	courseAssignments := api.GetAssignments(serverID)
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Assignments Due Today")
	p := message.NewPrinter(language.English)
	body := ""
	for _, course := range courseAssignments.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil {
			continue
		}
		assignmentsExist := false
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if (assignment.DueAt.Unix() < time.Now().AddDate(0, 0, 0).Unix()) || ((assignment.DueAt.Unix() - time.Now().AddDate(0, 0, 0).Unix()) > 57600) {
				continue
			}
			if assignmentsExist == false {
				body += p.Sprintf("__**%s**__\n", course.CourseName[5:])
				valid, assignmentsExist = true, true
			}
			days := int(time.Until(assignment.DueAt).Hours() / 24)
			hours := int(time.Until(assignment.DueAt).Hours() - float64(int(days*24)))
			minutes := int(time.Until(assignment.DueAt).Minutes() - float64(int(days*24*60)+int(hours*60)))
			body += p.Sprintf("%.0f Marks\n", assignment.PointsPossible)
			body += p.Sprintf("[%s](%s)\n", assignment.Name, assignment.HTMLURL)
			body += p.Sprintf("**%d Days, ", days)
			body += p.Sprintf("%d Hours, ", hours)
			body += p.Sprintf("%d Minutes** 	|	", minutes)
			body += p.Sprintf("%s\n\n", (assignment.DueAt.Format("02 Jan 2006 15:04")))
		}
	}
	if valid {
		s.ChannelMessageSend(alertChannelID, "@here")
		emb.SetDescription(body)
		s.ChannelMessageSendEmbed(alertChannelID, emb.MessageEmbed)
	}
}

func ModuleList(s *discordgo.Session, m *discordgo.MessageCreate) {

	courseAssignments := api.GetStats(m.GuildID)
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle(fmt.Sprintf("%d Module List", time.Now().Year()))

	p := message.NewPrinter(language.English)
	body := ""

	for _, course := range courseAssignments.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil {
			continue
		}
		valid = true
		graded := false
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.ScoreStatistics.Mean != 0 {
				graded = true
				break
			}
		}
		body += p.Sprintf("**%s**\n", course.CourseName[5:])
		body += p.Sprintf("[Canvas]("+viper.GetString("canvas.domain")+"/courses/%s) | ", course.ID)
		body += p.Sprintf("[UCC](https://www.ucc.ie/admin/registrar/modules/?mod=%s) | ", course.CourseCode[5:])
		switch graded {
		case true:
			body += "Stats: ✓\n\n"
		case false:
			body += "Stats: ✗\n\n"
		}
	}
	if valid {
		emb.Description = body
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "> **Error getting module data**")
	}
}

func CourseStats(s *discordgo.Session, m *discordgo.MessageCreate) {

	totalMax, totalMean := 0.0, 0.0
	cm, slug := extractCommand(m.Content)
	log.Println(slug)
	if cm == slug {
		s.ChannelMessageSend(m.ChannelID, "> **Please enter a valid module code**")
		return
	}
	slug = strings.ToUpper(strings.Split(slug, " ")[1])
	courseAssignments := api.GetStats(m.GuildID)

	valid := false
	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Available Grade Statistics: " + slug)

	body := "```\n"

	for _, course := range courseAssignments.Data.AllCourses {
		if (len(course.Term.Name) > 8) || course.EnrollmentsConnection.Nodes == nil || course.CourseCode[len(course.CourseCode)-6:] != slug {
			continue
		}
		for _, assignment := range course.AssignmentsConnection.Nodes {
			if assignment.ScoreStatistics.Max == 0 {
				continue
			}
			valid = true
			assName := ""
			if len(assignment.Name) >= 32 {
				assName = assignment.Name[:32] + "..."
			} else {
				assName = assignment.Name
			}
			body += assName + fmt.Sprintf(" (%.0f Marks)", assignment.PointsPossible) + ":\n--------------------------------------\n"

			body += "	" + fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Max)) + "	|	"
			body += fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Mean)) + "	|	"
			body += fmt.Sprintf("%.2f", (assignment.ScoreStatistics.Min)) + "	"
			body += "\n\n"
			totalMax += assignment.PointsPossible
			totalMean += assignment.ScoreStatistics.Mean
		}
	}
	if valid {
		body += "```"

		msg := "```\n"
		msg += "Total Mark Average (so far):\n"
		msg += "--------------------------------------\n"
		msg += "	" + fmt.Sprintf("%.2f / %.2f", (totalMean), (totalMax)) + "	|	"
		msg += fmt.Sprintf("%.2f", ((totalMean/totalMax)*(100/1))) + "%"
		msg += "\n```\n"
		msg += body

		emb.Description = msg

		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		s.ChannelMessageSend(m.ChannelID, "> **No module data found**")
	}
}

func CoordinatorInfo(s *discordgo.Session, m *discordgo.MessageCreate) {

	courseAssignments := api.GetAssignments(m.GuildID)
	valid := false

	emb := embed.NewEmbed()

	emb.SetColor(0xab0df9)

	emb.SetTitle("Module Coordinator Info")

	p := message.NewPrinter(language.English)
	for _, course := range courseAssignments.Data.AllCourses {
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
		s.ChannelMessageSend(m.ChannelID, "> **Error getting module data**")
	}
}
