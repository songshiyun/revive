package test

import (
	"testing"

	"github.com/songshiyun/revive/lint"
	"github.com/songshiyun/revive/rule"
)

// Test that left and right side of Binary operators (only AND, OR) are swapable
func TestOptimizeOperandsOrder(t *testing.T) {
	testRule(t, "optimize-operands-order", &rule.OptimizeOperandsOrderRule{}, &lint.RuleConfig{})
}
