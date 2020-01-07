package main

import (
	"TP2/Sub"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	print("\nProgramme Start")
	subject := new(Sub.Subject)
	subject.Id = 1

	rpc.Register(subject)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	time.Sleep(time.Second * 20)

	print("Programme Finished")
}
