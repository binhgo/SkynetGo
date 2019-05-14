package main

import (
	"log"
	"net/rpc"

	"SkynetGo/jrpc/service"
)

func main() {
	// make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}

	// -----------------------------------------------
	// make arguments object
	args := &service.Args{
		A: 2,
		B: 3,
	}
	// this will store returned result
	var result service.Result
	// call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	// we got our result in result
	log.Printf("%d*%d=%d\n", args.A, args.B, result)

	// -----------------------------------------------
	postAgrs := &service.PostArgs{Page: 1, UserID: 1}
	var postResult service.PostResult
	client.Call("PostRPC.GetPosts", postAgrs, &postResult)
	if err != nil {
		log.Fatalf("error in PostRPC", err)
	}

	for _, p := range postResult.Posts {
		log.Printf("%s\n", p)
	}

}
