package Obs

import "fmt"

type Observer struct {
	Id int
}
type RPCObserver interface {
 Notify(msg NotifyMsg, ack *bool) error
}
type NotifyMsg struct {
	 SubjectId int
	 Msg string
}

func (obs *Observer) Notify(msg NotifyMsg, ack *bool) error{
		fmt.Printf("Notify : %v \nMessage : %v\n",msg.SubjectId, msg.Msg)
	return nil
}

