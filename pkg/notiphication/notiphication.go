package notiphication

import (
	"log"
	"os/exec"
	"strings"
)

type Notiphication struct {
	Title         string
	Link          string
	Actions       []string
	DropdownLabel string
}

func (n Notiphication) buildCommand() ([]string, error) {
	command := []string{"-message", n.Title}
	if len(n.Actions) > 0 {
		command = append(command, "-actions", strings.Join(n.Actions, ","))
	}
	if n.DropdownLabel != "" {
		command = append(command, "-dropdownLabel", n.DropdownLabel)
	}
	return command, nil
}

func (n Notiphication) Push() (string, error) {

	command, err := n.buildCommand()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("alerter", command...)
	response, err := cmd.Output()
	if err != nil {

		log.Fatal(err)
	}
	return string(response), nil
}
