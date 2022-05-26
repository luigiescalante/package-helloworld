package helloworld

import (
	message "github.com/luigiescalante/package-helloworld/hellomessage"
)

func HelloWorld() *message.HelloMessage {
	return &message.HelloMessage{}
}
