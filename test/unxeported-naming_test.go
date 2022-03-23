package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

func TestUnexportednaming(t *testing.T) {
	testRule(t, "unexported-naming", &rule.UnexportedNamingRule{})
}
