package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type RPCSubject interface {
	Subscribe(o Observer, ack *bool) error
	Unsubscribe(o Observer, ack *bool) error
	Publish(msg string, ack *bool) error
}

type RPCObserver interface {
	Notify(msg NotifyMsg, ack *bool) error
}

type Subject struct {
	id int
	observers *list.List
}

type Observer struct {
	id int
}

type NotifyMsg struct {
	SubjectId int
	Msg string
}

func (s *Subject) Subscribe(o Observer, ack *bool) error {
	for e := s.observers.Front(); e != nil; e = e.Next() {
		if e.Value.(Observer).id == o.id {
			return errors.New("Already suscribed!")
		}
	}

	s.observers.PushBack(o)

	return nil
}

func (s *Subject) Unsubscribe(o Observer, ack *bool) error {
	for e := s.observers.Front(); e != nil; e = e.Next() {
		if e.Value.(Observer).id == o.id {
			s.observers.Remove(e)
			return nil
		}
	}
	return errors.New("Not suscribed!")
}

func (s *Subject) Publish(msg string, ack *bool) error {
	for e := s.observers.Front(); e != nil; e = e.Next() {
		o := e.Value.(Observer)
		var res bool
		var err = (&o).Notify(NotifyMsg{s.id, msg}, &res)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}

	return nil
}

func (o *Observer) Notify(msg NotifyMsg, ack *bool) error {
	fmt.Printf("Observer.id: %d\nSubject.id: %d\nMessage: %s\n", o.id, msg.SubjectId, msg.Msg)

	return nil
}

func main_server() {
	l := list.New()
	s := &Subject{1, l}
	// o_1 := Observer{1}
	// o_2 := Observer{2}
	// o_3 := Observer{3}


	rpc.Register(s)
	rpc.HandleHTTP()

	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(lis, nil)

	/*
		var res bool
		s.Subscribe(o_1, &res)
		s.Subscribe(o_2, &res)
		s.Subscribe(o_3, &res)

		s.Publish("Hello", &res)
		s.Publish("World", &res)
		s.Publish("🙂", &res)

		s.Unsubscribe(o_1, &res)
		s.Unsubscribe(o_2, &res)
		s.Unsubscribe(o_3, &res)
	*/
}