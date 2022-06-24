package uuids

import (
	"testing"

	guuid "github.com/google/uuid"
)

func TestGoogleUUID(t *testing.T) {
	i := 999
	for i > 0 {
		uuid := guuid.New().String()
		t.Log(uuid, len(uuid))
		i--
	}
}
