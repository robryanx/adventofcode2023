package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1":  "55607",
	"1-2":  "55291",
	"2-1":  "2771",
	"2-2":  "70924",
	"3-1":  "539713",
	"3-2":  "84159075",
	"4-1":  "23941",
	"4-2":  "5571760",
	"5-1":  "218513636",
	"5-2":  "81956384",
	"6-1":  "1312850",
	"6-2":  "36749103",
	"7-1":  "248179786",
	"7-2":  "247885995",
	"8-1":  "18113",
	"8-2":  "12315788159977",
	"9-1":  "1798691765",
	"9-2":  "1104",
	"10-1": "6864",
	"10-2": "349",
	"11-1": "10885634",
	"11-2": "707505470642",
	"12-1": "7753",
	"12-2": "280382734828319",
	"13-1": "28895",
	"13-2": "31603",
	"14-1": "109466",
	"14-2": "94585",
	"15-1": "507291",
	"15-2": "296921",
	"16-1": "6361",
	"16-2": "6701",
	"18-1": "76387",
	"19-1": "472630",
	"20-1": "980457412",
	"20-2": "232774988886497",
	"21-1": "3699",
	"22-1": "457",
	"22-2": "79122",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		expect := expect
		day := day
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", fmt.Sprintf("days/%s/main.go", day))
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(output)
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
