package uuid

import "testing"

func TestNewUUID(t *testing.T) {
	uniqID := NewUUID()
	t.Logf(uniqID)
}
