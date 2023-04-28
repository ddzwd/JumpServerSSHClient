package server_test

import (
	"testing"

	"JumpServerSSHClient/server"
)

func TestMFAGenerate(t *testing.T) {
	cases := []struct {
		secret   string
		timesmap uint64
		code     string
	}{
		{"K6UX6C54YBGVYNQD", 1680699691, "010629"},
		{"K6UX6C54YBGVYNQD", 1680799691, "225975"},
		{"K6UX6C54YBGVYNQK", 1680799691, "370376"},
		{"", 1680799691, "070026"},
	}

	for _, c := range cases {
		if server.MFAHmacSha1(c.secret, c.timesmap) != c.code {
			t.Errorf("secret: %s, timesmap: %d, code: %s", c.secret, c.timesmap, c.code)
		}
	}
}
