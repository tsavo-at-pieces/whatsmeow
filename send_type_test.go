package whatsmeow

import (
	"testing"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func TestGetTypeFromMessage(t *testing.T) {
	for _, tc := range []struct {
		name string
		msg  *waE2E.Message
		want string
	}{
		{"text", &waE2E.Message{Conversation: proto.String("hi")}, "text"},
		{"poll", &waE2E.Message{PollCreationMessage: &waE2E.PollCreationMessage{Name: proto.String("q")}}, "poll"},
		{"event", &waE2E.Message{EventMessage: &waE2E.EventMessage{Name: proto.String("e")}}, "event"},
		{"reaction", &waE2E.Message{ReactionMessage: &waE2E.ReactionMessage{Text: proto.String("x")}}, "reaction"},
	} {
		if got := getTypeFromMessage(tc.msg); got != tc.want {
			t.Errorf("%s: getTypeFromMessage = %q, want %q", tc.name, got, tc.want)
		}
	}
}
