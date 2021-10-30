package gogreeting

import (
	"errors"
	"fmt"
)

type Greeter interface {
	Hello(name string) (string, error)
}

type MyGreeter struct {
	Name string
}

func (g MyGreeter) Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("zero length name not allowed")
	}
	if name == "Steve" {
		return "", errors.New("name not allowed (Steve)")
	}
	return fmt.Sprintf("Hello %s", name), nil
}
