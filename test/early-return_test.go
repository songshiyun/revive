package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early-return", &rule.EarlyReturnRule{})
}
