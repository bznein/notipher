package notiphication

import (
	"log"
	"os/exec"
	"strings"
)

type Notiphication struct {
	Title         string
	Link          string
	Actions       Actions
	DropdownLabel string
}

const (
	CLICKED = "@CONTENTCLICKED"
	CLOSED  = "@CLOSED"
)

func (n Notiphication) buildCommand() ([]string, error) {
	command := []string{"-message", n.Title}
	if len(n.Actions) > 0 {
		command = append(command, "-actions", strings.Join(n.Actions.Keys(), ","))
	}
	if n.DropdownLabel != "" {
		command = append(command, "-dropdownLabel", n.DropdownLabel)
	}
	return command, nil
}

func (n Notiphication) SyncPush() {

	command, err := n.buildCommand()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("alerter", command...)
	response, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(response)
	switch responseString {
	case CLICKED:
		if n.Link != "" {
			exec.Command("open", n.Link).Start()
		}
	case CLOSED:
		return
	default:
		if action, ok := n.Actions[responseString]; ok {
			action()
		}
	}
}
