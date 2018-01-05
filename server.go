package main

import (
	"net/rpc"
	"net"
	"fmt"
	"net/http"
	//"time"
)

const (
	SAY = "CHello"
	PORT = ":10012"
	PROXY = "tcp"
)

type Hello struct {
	StrHello string
}

type ResStr struct {
	StrHi string
}

type Say int

func (this *Say) SayHello(args *Hello, reply *string) error {
	*reply = SAY + args.StrHello
	return nil
}

func (this *Say) SayHi(args *Hello, reply *ResStr) error {
	reply.StrHi = SAY + args.StrHello
	return nil
}

/*func client() {
	client, err := rpc.DialHTTP(CPROXY, CPORT)
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var args = CHello{" Test rpc",}
	var reply string
	err = client.Call("Say.SayHello", args, &reply)
	if err != nil {
		fmt.Println("arith error:", err)
	}
	fmt.Println(reply)
	args = CHello{" Go invokes the function asynchronously",}
	res := new(CResStr)
	resCall  := client.Go("Say.SayHi",args,&res,nil) //异步调用
	replyCall := <-resCall.Done
	if replyCall.Error != nil {
		fmt.Println("arith error:", replyCall.Error)
	}
	fmt.Println(res.StrHi)
}*/


func main() {
	hi := new(Say)
	rpc.Register(hi)
	rpc.HandleHTTP()
	l, e := net.Listen(PROXY,PORT)
	if e != nil {
		fmt.Println("listen error:", e)
	}
	http.Serve(l, nil)
	//for{
	//	time.Sleep(5*time.Second)
	//	client()
	//}
}
