package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestUnreachableCode(t *testing.T) {
	testRule(t, "unreachable-code", &rule.UnreachableCodeRule{})
}
