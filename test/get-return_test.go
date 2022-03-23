package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
