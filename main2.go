package main

import (
	"TP2/Obs"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Obs.Observer{Id: 1}
	var response *bool
	var response2 *bool
	call1 := client.Call("Subject.Subscribe", args, &response)
	fmt.Printf("\n Response: ", *response)
	time.Sleep(time.Second * 10)
	call2 := client.Call("Subject.Unsubscribe", args, &response2)
	fmt.Printf("\n Response2: ", *response)

	if call1 != nil {
		fmt.Printf("Erro : %v", call1)
	}
	if call2 != nil {
		fmt.Printf("Error : %v", call2)
	}

}
