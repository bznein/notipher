package main

import (
	"fmt"
	"log"

	"github.com/bznein/notipher/pkg/notiphication"
)

func main() {
	notip := notiphication.Notiphication{}

	actions := notiphication.Actions{}
	actions["action1"] = func() { fmt.Println("Clicked action1") }
	actions["action2"] = func() { fmt.Println("Clicked action2") }
	notip.Title = "Title"
	notip.Message = "Message"
	notip.Subtitle = "Subtitle"
	notip.Close = "CloseLabel"
	notip.CloseFunc = func() { fmt.Println("Close func") }
	notip.Timeout = 100
	notip.TimeoutFunc = func() { fmt.Println("Timeout func") }
	//notip.Reply = "Reply"
	notip.Actions = actions
	notip.Link = "http://www.google.com"
	notip.DropdownLabel = "label"
	ret, err := notip.SyncPush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}
