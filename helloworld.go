package helloworld

import (
	message "github.com/luigiescalante/go-package-helloworld/hellomessage"
)

func HelloWorld() *message.HelloMessage {
	return &message.HelloMessage{}
}
