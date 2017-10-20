package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bkono/msgme/proto/msgme"
	"github.com/go-log/log"
	"github.com/micro/protobuf/ptypes"
	"golang.org/x/net/context"
)

// MsgMe implements the MsgMeHandler service interface.
type MsgMe struct{}

// Send relays a message to all current listeners.
func (m *MsgMe) Send(ctx context.Context, req *msgme.Message, rsp *msgme.SendResponse) error {
	log.Log("Received MsgMe.Send call")
	return nil
}

// Listen sets up a stream for receiving all future messages.
func (m *MsgMe) Listen(ctx context.Context, req *msgme.ListenRequest, stream msgme.MsgMe_ListenStream) error {
	log.Log("Received MsgMe.Listen call")

	for i := 0; i < 5; i++ {
		log.Logf("sending a message")
		if err := stream.Send(&msgme.Message{
			From:    "random",
			Content: fmt.Sprintf("some content %v", i),
			SentAt:  ptypes.TimestampNow(),
		}); err != nil {
			return err
		}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

	return nil
}
