package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
}
