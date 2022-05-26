package hellomessage

import "strings"

type HelloMessage struct {
	message string
}

func (hello HelloMessage) SayHello(name string) string {
	return "Hello from private repo to " + hello.formatName(name)
}

func (hello HelloMessage) formatName(name string) string {
	return strings.ToUpper(name)
}
