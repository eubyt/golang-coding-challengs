package dns

import (
	"bytes"
	"testing"
)

func TestResourceRecord(t *testing.T) {
	t.Run("should encode name and convert resource record to bytes", func(t *testing.T) {
		record := NewResourceRecord("www.google.com", 1, 1, 60, 4, []byte{216, 58, 211, 196})
		result := record.ToBytes()
		expected := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1, 0, 0, 0, 60, 0, 4, 216, 58, 211, 196}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("should parse resource record bytes", func(t *testing.T) {
		data := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1, 0, 0, 0, 60, 0, 4, 216, 58, 211, 196}

		record, _, err := BytesToResourceRecord(data)
		if err != nil {
			t.Fatalf("BytesToResourceRecord returned unexpected error: %v", err)
		}

		expectedName := "www.google.com"
		if record.Name != expectedName {
			t.Fatalf("Expected Name %q, got %q", expectedName, record.Name)
		}

		if record.Type != 1 {
			t.Fatalf("Expected Type 1, got %d", record.Type)
		}

		if record.Class != 1 {
			t.Fatalf("Expected Class 1, got %d", record.Class)
		}

		if record.TTL != 60 {
			t.Fatalf("Expected TTL 60, got %d", record.TTL)
		}

		if record.RDLength != 4 {
			t.Fatalf("Expected RDLength 4, got %d", record.RDLength)
		}

		if !bytes.Equal(record.RData, []byte{216, 58, 211, 196}) {
			t.Fatalf("Expected RData %v, got %v", []byte{216, 58, 211, 196}, record.RData)
		}
	})

	t.Run("should return error for malformed resource record without name terminator", func(t *testing.T) {
		data := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm'}

		if _, _, err := BytesToResourceRecord(data); err == nil {
			t.Fatal("Expected error for malformed resource record without name terminator")
		}
	})
}
