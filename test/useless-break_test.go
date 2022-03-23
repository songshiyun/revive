package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

// UselessBreak rule.
func TestUselessBreak(t *testing.T) {
	testRule(t, "useless-break", &rule.UselessBreak{})
}
