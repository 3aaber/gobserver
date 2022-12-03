# Golang Pub/Sub (Observer Pattern) Library

## Description

Golang Pub/Sub (Observer Pattern) implementation with support of channel and sync/async publish

## Example

```golang

import (
    "fmt"

    observer "github.com/sabermesgari/gobserver"
)

type Subject struct {
    ID      int
    Message string
}

// Notify
func (t *Subject) Notify(m interface{}) {
    t.Message = m.(string)
}

func main() {
    publisher := observer.NewPublisher()
    subject1 := &Subject{1, "Subject1"}
    publisher.AddSubject(subject1, "channel_1")
    publisher.NotifySubjectsSync("Message-1", "channel_1")
    fmt.Println(subject1.Message)
}

```
