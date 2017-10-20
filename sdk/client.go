package msgmesdk

import (
	"io"

	"github.com/bkono/msgme/proto/msgme"
	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/selector/cache"
	"github.com/micro/protobuf/ptypes"
	"golang.org/x/net/context"
)

type MessageCallback interface {
	OnMessage(from, message string, sentAt int64)
}

type Client struct {
	cl msgme.MsgMeClient
	cb MessageCallback
}

func NewClient() *Client {
	cl := msgme.NewMsgMeClient("go.micro.srv.msgme", client.NewClient(
		client.Selector(cache.NewSelector(
			selector.Registry(mdns.NewRegistry())))))

	return &Client{cl: cl}
}

func (c *Client) Send(from, message string) error {
	_, err := c.cl.Send(context.Background(), &msgme.Message{
		From:    from,
		Content: message,
		SentAt:  ptypes.TimestampNow(),
	})
	return err
}

func (c *Client) Listen(cb MessageCallback) {
	c.cb = cb
	go c.startListen()
}
func (c *Client) startListen() {
	stream, err := c.cl.Listen(context.Background(), &msgme.ListenRequest{})
	if err != nil {
		log.Logf("err starting stream: %v\n", err)
		return
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Log("eof reached. exiting gracefully\n")
			break
		}

		if err != nil {
			log.Logf("err in stream: %v\n", err)
			break // exit ungracefully?
		}

		log.Logf("message received: %v\n", msg)
		c.cb.OnMessage(msg.From, msg.Content, msg.SentAt.Seconds)
	}
}
