package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Grumpy  bool
	Message Message
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{
		Grumpy:  grumpy,
		Message: m,
	}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("go away!")
	}
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("cloud not creat event : event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage("hi there!")
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		log.Fatal(err)
	}
	event.Start()
}
