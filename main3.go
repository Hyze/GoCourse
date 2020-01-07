package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := "Vous avez été notfié"
	var response *bool

	call := client.Call("Subject.Publish", args, &response)
	fmt.Printf("\n Response: ", *response)

	if call != nil {
		fmt.Printf("Error : %v", call)
	}

}
