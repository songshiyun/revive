package test

import (
	"testing"

	"github.com/songshiyun/revive/rule"
)

// TestTimeEqual rule.
func TestTimeEqual(t *testing.T) {
	testRule(t, "time-equal", &rule.TimeEqualRule{})
}
