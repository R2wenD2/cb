package cb

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	if got != 2 {
		t.Errorf("Add(1,1) = %d; want 2", got)
	}
}
