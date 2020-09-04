package notiphication

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Notiphication struct {
	Title         string
	Message       string
	Subtitle      string
	Link          string
	Actions       Actions
	DropdownLabel string
	Reply         string
	Close         string
	CloseFunc     func()
	Timeout       uint64
	TimeoutFunc   func()
}

type executionResult struct {
	RetVal string
	Err    error
}

const (
	CLICKED = "@CONTENTCLICKED"
	CLOSED  = "@CLOSED"
	TIMEOUT = "@TIMEOUT"
)

func (n Notiphication) validate() error {
	if n.Reply != "" && len(n.Actions) != 0 {
		return fmt.Errorf("Can not specify both Reply and Actions")
	}
	if n.DropdownLabel != "" && len(n.Actions) == 0 {
		return fmt.Errorf("Can not specify DropdownLabel is no Actions are specified")
	}
	return nil
}

func (n Notiphication) buildCommand() []string {
	command := []string{"-message", n.Message}
	if len(n.Actions) > 0 {
		command = append(command, "-actions", strings.Join(n.Actions.Keys(), ","))
	}
	if n.DropdownLabel != "" {
		command = append(command, "-dropdownLabel", n.DropdownLabel)
	}
	if n.Reply != "" {
		command = append(command, "-reply", n.Reply)
	}
	if n.Subtitle != "" {
		command = append(command, "-subtitle", n.Subtitle)
	}
	if n.Title != "" {
		command = append(command, "-title", n.Title)
	}
	if n.Close != "" {
		command = append(command, "-closeLabel", n.Close)
	}
	if n.Timeout != 0 {
		command = append(command, "-timeout", strconv.FormatUint(n.Timeout, 10))
	}
	return command
}

func (n Notiphication) send(c chan executionResult) executionResult {
	if err := n.validate(); err != nil {
		select {
		case c <- executionResult{"", err}:
		default:
			return executionResult{"", err}
		}
	}
	command := n.buildCommand()
	cmd := exec.Command("alerter", command...)
	response, err := cmd.Output()
	if err != nil {
		select {
		case c <- executionResult{"", err}:
		default:
			return executionResult{"", err}
		}
	}
	responseString := string(response)
	switch responseString {
	case CLICKED:
		if n.Link != "" {
			exec.Command("open", n.Link).Start()
		}
	case CLOSED:
	case n.Close:
		n.CloseFunc()
		select {
		case c <- executionResult{"", nil}:
		default:
			return executionResult{"", nil}
		}
	case TIMEOUT:
		n.TimeoutFunc()
		select {
		case c <- executionResult{"", nil}:
		default:
			return executionResult{"", nil}
		}
	default:
		if action, ok := n.Actions[responseString]; ok {
			action()
		} else {
			fmt.Println("sending")
			select {
			case c <- executionResult{responseString, nil}:
			default:
				return executionResult{responseString, nil}
			}
		}
	}
	select {
	case c <- executionResult{"", nil}:
	default:
	}
	return executionResult{"", nil}
}

func (n Notiphication) SyncPush() (string, error) {
	c := make(chan executionResult)
	res := n.send(c)
	return res.RetVal, res.Err

}

func (n Notiphication) AsyncPush() (string, error) {
	c := make(chan executionResult)
	go n.send(c)
	res := <-c
	return res.RetVal, res.Err
}
