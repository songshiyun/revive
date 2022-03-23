package test

import (
	"testing"

	"github.com/songshiyun/revive/lint"
	"github.com/songshiyun/revive/rule"
)

func TestNestedStructs(t *testing.T) {
	testRule(t, "nested-structs", &rule.NestedStructs{}, &lint.RuleConfig{})
}
