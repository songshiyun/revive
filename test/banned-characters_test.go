package test

import (
	"testing"

	"github.com/songshiyun/revive/lint"
	"github.com/songshiyun/revive/rule"
)

// Test banned characters in a const, var and func names.
// One banned character is in the comment and should not be checked.
// One banned character from the list is not present in the fixture file.
func TestBannedCharacters(t *testing.T) {
	testRule(t, "banned-characters", &rule.BannedCharsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"Ω", "Σ", "σ", "1"},
	})
}
