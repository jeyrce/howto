package pkg

import (
	"testing"
	"time"

	sf "github.com/sony/sonyflake"
)

func TestUUID(t *testing.T) {
	uuid := sf.NewSonyflake(sf.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			return 7, nil
		},
		CheckMachineID: func(u uint16) bool {
			return u == 7
		},
	})
	for i := 0; i <= 999; i++ {
		id, err := uuid.NextID()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id, sf.Decompose(id))
	}
}
