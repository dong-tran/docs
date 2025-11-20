package behavioral

import "fmt"

// Mediator Pattern - Reduces coupling between components by making them communicate through a mediator.

type ChatMediator interface {
	SendMessage(message string, user User)
	AddUser(user User)
}

type User interface {
	Send(message string)
	Receive(message string)
	GetName() string
}

type ChatRoom struct {
	users []User
}

func (c *ChatRoom) SendMessage(message string, user User) {
	for _, u := range c.users {
		if u.GetName() != user.GetName() {
			u.Receive(fmt.Sprintf("[%s]: %s", user.GetName(), message))
		}
	}
}

func (c *ChatRoom) AddUser(user User) {
	c.users = append(c.users, user)
	fmt.Printf("%s joined the chat\n", user.GetName())
}

type ChatUser struct {
	name     string
	mediator ChatMediator
}

func NewChatUser(name string, mediator ChatMediator) *ChatUser {
	return &ChatUser{name: name, mediator: mediator}
}

func (u *ChatUser) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}

func (u *ChatUser) Receive(message string) {
	fmt.Printf("%s receives: %s\n", u.name, message)
}

func (u *ChatUser) GetName() string {
	return u.name
}

func DemoMediator() {
	fmt.Println("=== Mediator Pattern Demo ===\n")
	chatRoom := &ChatRoom{}
	alice := NewChatUser("Alice", chatRoom)
	bob := NewChatUser("Bob", chatRoom)
	charlie := NewChatUser("Charlie", chatRoom)
	chatRoom.AddUser(alice)
	chatRoom.AddUser(bob)
	chatRoom.AddUser(charlie)
	fmt.Println()
	alice.Send("Hello everyone!")
	fmt.Println()
	bob.Send("Hi Alice!")
}
