package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

// TestModifiesParam rule.
func TestModifiesParam(t *testing.T) {
	testRule(t, "modifies-param", &rule.ModifiesParamRule{})
}
