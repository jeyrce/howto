package cpu

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func TestCPU(t *testing.T) {
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(percent)
}
