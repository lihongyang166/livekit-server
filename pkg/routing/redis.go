package routing

import (
	"context"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/livekit-server/proto/livekit"
)

const (
	// hash of node_id => Node proto
	NodesKey = "nodes"

	// hash of room_name => node_id
	NodeRoomKey = "room_node_map"
)

var redisCtx = context.Background()

// location of the participant's RTC connection, hash
func participantRTCKey(participantId string) string {
	return "participant_rtc:" + participantId
}

// location of the participant's Signal connection, hash
func participantSignalKey(participantId string) string {
	return "participant_signal:" + participantId
}

func nodeChannel(nodeId string) string {
	return "node_channel:" + nodeId
}

func publishRouterMessage(rc *redis.Client, nodeId string, participantId string, msg proto.Message) error {
	rm := &livekit.RouterMessage{
		ParticipantId: participantId,
	}
	switch o := msg.(type) {
	case *livekit.StartSession:
		rm.Message = &livekit.RouterMessage_StartSession{
			StartSession: o,
		}
	case *livekit.SignalRequest:
		rm.Message = &livekit.RouterMessage_Request{
			Request: o,
		}
	case *livekit.SignalResponse:
		rm.Message = &livekit.RouterMessage_Response{
			Response: o,
		}
	case *livekit.EndSession:
		rm.Message = &livekit.RouterMessage_EndSession{
			EndSession: o,
		}
	default:
		return errInvalidRouterMessage
	}
	data, err := proto.Marshal(rm)
	if err != nil {
		return err
	}
	return rc.Publish(redisCtx, nodeChannel(nodeId), data).Err()
}

type RedisSink struct {
	rc            *redis.Client
	nodeId        string
	participantId string
	channel       string
	onClose       func()
}

func (s *RedisSink) WriteMessage(msg proto.Message) error {
	return publishRouterMessage(s.rc, s.nodeId, s.participantId, msg)
}

func (s *RedisSink) Close() {
	publishRouterMessage(s.rc, s.nodeId, s.participantId, &livekit.EndSession{})
	if s.onClose != nil {
		s.onClose()
	}
}

func (s *RedisSink) OnClose(f func()) {
	s.onClose = f
}
