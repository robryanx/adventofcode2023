package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1": "55607",
	"1-2": "55291",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		cmd := exec.Command(fmt.Sprintf("bin/%s", day))
		out, err := cmd.CombinedOutput()

		assert.NoError(t, err)
		assert.Equal(t, expect, strings.TrimRight(string(out), "\n"), fmt.Sprintf("Day %s", day))
	}
}
