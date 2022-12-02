package main

import (
	"fmt"

	observer "github.com/sabermesgari/gobserver"
)

type Subject struct {
	ID      int
	Message string
}

// Notify Notify
func (t *Subject) Notify(m interface{}) {
	t.Message = m.(string)
}

func main() {
	publisher := observer.NewPublisher()

	subject1 := &Subject{1, "Subject1"}
	subject2 := &Subject{2, "Subject2"}
	subject3 := &Subject{3, "Subject3"}

	publisher.AddSubject(subject1, "channel_1")
	publisher.AddSubject(subject2, "channel_2")
	publisher.AddSubject(subject3, "channel_3")

	publisher.NotifySubjectsSync("Message-1", "channel_1")
	fmt.Println(subject1.Message)

	publisher.NotifySubjectsSync("Message-2", "channel_2")
	fmt.Println(subject2.Message)

	publisher.NotifySubjectsSync("Message-3", "channel_3")
	fmt.Println(subject3.Message)

}
