package dns

import (
	"bytes"
	"testing"
)

func TestBytesHeader(t *testing.T) {
	t.Run("should convert header to bytes", func(t *testing.T) {
		headerFlag := NewHeaderFlag(false, 0, false, false, true, false, 0, 0).ToUint16()
		header := NewHeader(22, headerFlag, 1, 0, 0, 0)
		result := header.ToBytes()
		expected := []byte{0, 22, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("should convert bytes to header", func(t *testing.T) {
		data := []byte{0, 22, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		result, err := BytesToHeader(data)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := &Header{
			ID:      22,
			Flags:   NewHeaderFlag(false, 0, false, false, true, false, 0, 0).ToUint16(),
			QDCount: 1,
			ANCount: 0,
			NSCount: 0,
			ARCount: 0,
		}
		if *result != *expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
