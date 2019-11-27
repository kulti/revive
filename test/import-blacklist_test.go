package test

import (
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

func TestImportsBlacklist(t *testing.T) {
	args := []interface{}{"crypto/md5", "crypto/sha1"}

	testRule(t, "imports-blacklist", &rule.ImportsBlacklistRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}

func BenchmarkImportsBlacklist(b *testing.B) {
	args := []interface{}{"crypto/md5", "crypto/sha1"}
	benchRule(b, "imports-blacklist", &rule.ImportsBlacklistRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
