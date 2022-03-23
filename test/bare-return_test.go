package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestBareReturn(t *testing.T) {
	testRule(t, "bare-return", &rule.BareReturnRule{})
}
