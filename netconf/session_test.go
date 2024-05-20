package netconf

import (
	"strings"
	"testing"
)

func TestV11SeparatorRemoval(t *testing.T) {
	preData := []string{
		"<rpc",
		"#7",
		"></rpc>",
	}
	postData := "<rpc></rpc>"
	result := stripV11Separators(strings.Join(preData, "\n"))
	if result != postData {
		t.Errorf("expected \"%s\", got \"%s", postData, result)

	}
}
