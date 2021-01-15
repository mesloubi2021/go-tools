package quickfix

import (
	"testing"

	"honnef.co/go/tools/lint/testutil"
)

func TestAll(t *testing.T) {
	checks := map[string][]testutil.Test{
		"QF1000": {{Dir: "CheckStringsIndexByte"}},
		"QF1001": {{Dir: "CheckDeMorgan"}},
		"QF1002": {{Dir: "CheckTaglessSwitch"}},
	}

	testutil.Run(t, Analyzers, checks)
}
