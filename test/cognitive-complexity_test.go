package test

import (
	"testing"

	"github.com/songshiyun/revive/lint"
	"github.com/songshiyun/revive/rule"
)

func TestCognitiveComplexity(t *testing.T) {
	testRule(t, "cognitive-complexity", &rule.CognitiveComplexityRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(0)},
	})
}
