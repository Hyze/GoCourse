package Sub

import (
	"TP2/Obs"
	"fmt"
)

type Subject struct {
	Id        int
	observers []Obs.Observer
}

type RPCSubject interface {
	Subscribe(o Obs.Observer, ack *bool) error
	Unsubscribe(o Obs.Observer, ack *bool) error
	Publish(msg string, ack *bool) error
}

func (sub *Subject) Subscribe(o Obs.Observer, ack *bool) error {
	sub.observers = append(sub.observers,o )
	*ack = true
	fmt.Printf("\nSubscribed -> obs : %v\n",o )
	return nil
}

func (sub *Subject) Unsubscribe(o Obs.Observer, ack *bool) error {
	for  i := range sub.observers {
		if sub.observers[i].Id == o.Id {
			sub.observers = append(sub.observers[:i])
			*ack = true
			fmt.Printf("Unsubscribed -> obs : %v\n",o )
		}
	}
	return nil
}

func (sub *Subject) Publish(msg string, ack *bool) error {
	for i := range sub.observers{
		*ack =true
		sub.observers[i].Notify(Obs.NotifyMsg{SubjectId: sub.Id, Msg: msg}, ack)
	}
	return nil
}
