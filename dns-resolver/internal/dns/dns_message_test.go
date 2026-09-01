package dns

import (
	"bytes"
	"testing"
)

func TestDNSMessage(t *testing.T) {
	t.Run("should convert DNS message to bytes", func(t *testing.T) {
		headerFlag := NewHeaderFlag(false, 0, false, false, true, false, 0, 0).ToUint16()
		header := NewHeader(22, headerFlag, 1, 0, 0, 0)
		question := NewQuestion("www.google.com", 1, 1)
		DNSMessage := NewDNSMessage(*header, []Question{*question})
		result := DNSMessage.ToBytes()
		expected := []byte{0, 22, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0, 0, 1, 0, 1}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("should reject malformed dns message with insufficient bytes", func(t *testing.T) {
		if _, err := BytesToDNSMessage([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}); err == nil {
			t.Fatal("Expected error for malformed dns message shorter than 12 bytes")
		}
	})

	t.Run("should parse compressed names used in DNS responses", func(t *testing.T) {
		data := []byte{
			0x00, 0x01,
			0x81, 0x80,
			0x00, 0x01,
			0x00, 0x01,
			0x00, 0x00,
			0x00, 0x00,
			0x03, 'w', 'w', 'w',
			0x06, 'g', 'o', 'o', 'g', 'l', 'e',
			0x03, 'c', 'o', 'm', 0x00,
			0x00, 0x01,
			0x00, 0x01,
			0xc0, 0x0c,
			0x00, 0x01,
			0x00, 0x01,
			0x00, 0x00, 0x01, 0x1f,
			0x00, 0x04,
			0x8e, 0xfb, 0x99, 0x77,
		}

		msg, err := BytesToDNSMessage(data)
		if err != nil {
			t.Fatalf("BytesToDNSMessage returned unexpected error: %v", err)
		}

		if len(msg.Questions) != 1 {
			t.Fatalf("Expected 1 question, got %d", len(msg.Questions))
		}
		if msg.Questions[0].Qname != "www.google.com" {
			t.Fatalf("Expected question Qname to be decoded, got %q", msg.Questions[0].Qname)
		}
		if len(msg.Answers) != 1 {
			t.Fatalf("Expected 1 answer, got %d", len(msg.Answers))
		}
		if msg.Answers[0].Name != "www.google.com" {
			t.Fatalf("Expected answer name pointer to be resolved, got %q", msg.Answers[0].Name)
		}
		if msg.Answers[0].RData[0] != 0x8e || msg.Answers[0].RData[3] != 0x77 {
			t.Fatalf("Expected RData to decode correctly, got %v", msg.Answers[0].RData)
		}
	})
}
