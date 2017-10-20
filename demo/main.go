package main

import (
	"fmt"
	"io"
	"os"

	"github.com/bkono/msgme/proto/msgme"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/selector/cache"
	"github.com/micro/protobuf/ptypes"
	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	microcl := client.NewClient(
		client.Selector(cache.NewSelector(
			selector.Registry(mdns.NewRegistry()))))
	cl := msgme.NewMsgMeClient("go.micro.srv.msgme", microcl)
	ctx := context.Background()

	fmt.Println("proving it works with a simple send")
	rsp, err := cl.Send(ctx, &msgme.Message{
		From:    "cli",
		Content: "some content from cli",
		SentAt:  ptypes.TimestampNow(),
	})
	exitOnErr(err)
	fmt.Printf("rsp(%v)\n", rsp)

	fmt.Println("proving streaming works")
	stream, err := cl.Listen(ctx, &msgme.ListenRequest{})
	exitOnErr(err)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream reached, proceed happy")
			break
		}

		if err != nil {
			fmt.Println("err receiving from stream: ", err)
			break
		}

		fmt.Printf("from(%v) content(%v) sentat(%v)\n", msg.From, msg.Content, msg.SentAt)
	}

	fmt.Println("demo complete")
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println("error found: ", err)
		os.Exit(1)
	}
}
