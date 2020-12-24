package commands

import (
	"fmt"

	"github.com/ReeceDonovan/uni-bot/request"
	"github.com/bwmarrin/discordgo"
)

// TODO: Add functions for more bot commands

// Test command for now, sends basic string of modules/assignments from current term as message. Will eventually be an embed and not hardcoded termID

func TermAssignments(s *discordgo.Session, m *discordgo.MessageCreate) {
	CourseAssignment := request.QueryAssignments()
	msg := ""
	for _, course := range CourseAssignment.Data.AllCourses {
		if course.Term.ID != "44" || course.EnrollmentsConnection == nil {
			continue
		}
		msg += "\n" + course.CourseName + ":\n"
		for _, assignment := range course.AssignmentsConnection.Edges {
			msg += fmt.Sprintf("%s", assignment.Node.HTMLURL) + "\n"
		}
	}
	fmt.Println(msg)
	s.ChannelMessageSend(m.ChannelID, msg)
}
