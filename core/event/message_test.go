package event

import "testing"

func TestMessage_GetType(t *testing.T) {
	msg := Message{}
	msg.SetType(MessageType("foo"))
	if msg.GetType() != MessageType("foo") {
		t.Errorf("Unexpected message type. Got %s, but expected %s", MessageType("foo"), msg.GetType())
	}
}

func TestMessage_SetType(t *testing.T) {
	msg := Message{}
	msg.SetType(MessageType("foo"))
	if msg.GetType() != MessageType("foo") {
		t.Errorf("Unexpected message type. Got %s, but expected %s", MessageType("foo"), msg.GetType())
	}
}
