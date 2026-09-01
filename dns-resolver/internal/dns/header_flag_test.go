package dns

import (
	"testing"
)

func TestHeaderFlag(t *testing.T) {
	t.Run("should convert header flag to uint16", func(t *testing.T) {
		headerFlag := NewHeaderFlag(true, 0, false, false, true, false, 0, 0)
		result := headerFlag.ToUint16()
		expected := uint16(32768 | 256)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
