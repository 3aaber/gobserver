package gobserver

import "log"

// Observer inteface to implement Notify
type Observer interface {
	Notify(interface{})
}

// Publisher is struct to encapsulate Observer List data
type Publisher struct {
	ObserversList map[string][]Observer
}

// NewPublisher make a new publisher object
func NewPublisher() *Publisher {
	p := new(Publisher)
	p.ObserversList = make(map[string][]Observer)
	return p
}

// AddSubject add a observer implemented from Observer inteface
func (s *Publisher) AddSubject(o Observer, channels ...string) {
	if len(channels) == 0 {
		// if the channels list is empty, add observer to all existing channels
		for chanl := range s.ObserversList {
			s.ObserversList[chanl] = append(s.ObserversList[chanl], o)
		}
	} else {
		// if the channels list is not empty, only add observer to that channel list
		for _, chanl := range channels {
			s.ObserversList[chanl] = append(s.ObserversList[chanl], o)
		}
	}
}

// RemoveSubject remove a observer from list
func (s *Publisher) RemoveSubject(o Observer, channels ...string) {

	if len(channels) == 0 {
		// Remove observer from all channels
		for chanl := range s.ObserversList {
			s.removeObserverSingleChannel(o, chanl)
		}

	} else {
		for _, chanl := range channels {
			s.removeObserverSingleChannel(o, chanl)
		}
	}
}

// removeObserverSingleChannel observer from a single channel
func (s *Publisher) removeObserverSingleChannel(o Observer, channel string) {

	// Check if the channel exist
	if _, ok := s.ObserversList[channel]; ok {

		// find the observer from that channel list
		for indexToRemove, observer := range s.ObserversList[channel] {
			if observer == o {

				// delete the observer from channel list
				s.ObserversList[channel] = append(s.ObserversList[channel][:indexToRemove],
					s.ObserversList[channel][indexToRemove+1:]...)
				break
			}
		}

	}

}

// NotifySubjectsSync in sync mode for a specific channel
func (s *Publisher) NotifySubjectsSync(m interface{}, channel string) {
	for _, observer := range s.ObserversList[channel] {
		log.Println(channel, m)
		observer.Notify(m)
	}
}

// NotifySubjectsASync usin g goroutin  for a specific channel
func (s *Publisher) NotifySubjectsASync(m interface{}, channel string) {
	run := func(message interface{}) {
		for _, observer := range s.ObserversList[channel] {
			observer.Notify(m)
		}
	}
	go run(m)
}
