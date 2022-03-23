package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestWaitGroupByValue(t *testing.T) {
	testRule(t, "waitgroup-by-value", &rule.WaitGroupByValueRule{})
}
