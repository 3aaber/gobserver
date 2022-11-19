package observer

import (
	"log"
	"testing"
)

// MessageFormat is a struct for json
type MessageFormat struct {
	InstrumentID string `json:"instrument"`
	Action       string `json:"action"`
	DateTime     string `json:"datetime"`
}

type TestObserver struct {
	ID      int
	Message interface{}
}

func (t *TestObserver) GetData(method string, queryData interface{}) interface{} {
	return nil
}

// GetName Get Name of Observer
func (t *TestObserver) GetName() string {
	return ""
}

// Notify Notify
func (t *TestObserver) Notify(m interface{}) {
	log.Printf("Observer %d: message '%s' received \n", t.ID, m.(MessageFormat).InstrumentID)
	t.Message = m
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, "default"}
	testObserver2 := &TestObserver{2, "default"}
	testObserver3 := &TestObserver{3, "default"}

	publisher := NewPublisher()

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1, "default")
		publisher.AddObserver(testObserver2, "default")
		publisher.AddObserver(testObserver3, "default")

		if len(publisher.ObserversList["default"]) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserversList["default"]) != 2 {
			t.Errorf("The size of the observer list is not the "+
				"expected. 3 != %d\n", len(publisher.ObserversList["default"]))
		}

		for _, observer := range publisher.ObserversList["default"] {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {

		if len(publisher.ObserversList) == 0 {
			t.Errorf("The list is empty. Nothing to test\n")
		}

		for _, observer := range publisher.ObserversList["default"] {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != "default" {
				t.Errorf("The observer's Message field weren't"+
					" empty: %s\n", printObserver.Message)
			}
		}

		message := MessageFormat{}
		message.InstrumentID = "Hello World!"
		publisher.NotifyObserversSync(message, "default")

		for _, observer := range publisher.ObserversList["default"] {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != message {
				t.Errorf("Expected message on observer %d was "+
					"not expected: '%s' != '%s'\n", printObserver.ID,
					printObserver.Message, message.InstrumentID)
			}
		}
	})
}
