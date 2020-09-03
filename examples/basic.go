package main

import (
	"github.com/bznein/notipher/pkg/notiphication"
)

func main() {
	notiphication := notiphication.Notiphication{}

	notiphication.Title = "aaa"
	notiphication.Actions = []string{"action", "action2"}
	notiphication.DropdownLabel = "label"
	notiphication.Push()
}
