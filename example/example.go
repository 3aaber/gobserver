package main

import (
	"fmt"

	observer "github.com/sabermesgari/gobserver"
)

type SampleObserver struct {
	ID      int
	Message string
}

// Notify Notify
func (t *SampleObserver) Notify(m interface{}) {
	t.Message = m.(string)
}

func main() {
	publisher := observer.NewPublisher()

	sampleObserver1 := &SampleObserver{1, "Observer1"}
	sampleObserver2 := &SampleObserver{2, "Observer2"}
	sampleObserver3 := &SampleObserver{3, "Observer3"}

	publisher.AddObserver(sampleObserver1, "channel_1")
	publisher.AddObserver(sampleObserver2, "channel_2")
	publisher.AddObserver(sampleObserver3, "channel_3")

	publisher.NotifyObserversSync("Message-1", "channel_1")
	fmt.Println(sampleObserver1.Message)

	publisher.NotifyObserversSync("Message-2", "channel_2")
	fmt.Println(sampleObserver2.Message)

	publisher.NotifyObserversSync("Message-3", "channel_3")
	fmt.Println(sampleObserver3.Message)

}
