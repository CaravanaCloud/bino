package commands

import (
	"bytes"
	"strings"

	mentors "github.com/CaravanaCloud/bino/mentors"
)

var mentorshipCommand = command{
	CanRun: canRun,
	Run:    run,
}

var Mentorship = CommandWrapper{
	Command: mentorshipCommand,
}

func canRun(message string) bool {
	return strings.TrimSpace(message) == "mentorias"
}

func run(message string) (string, error) {
	var stringBuffer bytes.Buffer
	for _, mentor := range mentors.List() {
		stringBuffer.WriteString("- " + mentor.Name + "\n")
	}
	return stringBuffer.String(), nil
}
