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
	"2-1": "2771",
	"2-2": "70924",
	"3-1": "539713",
	"3-2": "84159075",
	"4-1": "23941",
	"4-2": "5571760",
	"5-1": "218513636",
	"5-2": "81956384",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		cmd := exec.Command(fmt.Sprintf("bin/%s", day))
		out, err := cmd.CombinedOutput()

		assert.NoError(t, err)
		assert.Equal(t, expect, strings.TrimRight(string(out), "\n"), fmt.Sprintf("Day %s", day))
	}
}
