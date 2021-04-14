package whatthecommit_test

import (
	"github.com/mnlwldr/whatthecommit"
	"testing"
)

func TestMessage(t *testing.T) {
	message := whatthecommit.Message()
	if "" == message {
		t.Error("Message is empty string")
	}
}
