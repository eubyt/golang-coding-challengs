package dns

import (
	"bytes"
	"testing"
)

func TestQuestion(t *testing.T) {
	t.Run("should encode to name and convert question to bytes", func(t *testing.T) {
		name := encodeName("www.google.com")
		if name != "\x03www\x06google\x03com\x00" {
			t.Errorf("Expected encoded name to be %v, but got %v", "\x03www\x06google\x03com\x00", name)
		}
	})

	t.Run("should convert question to bytes", func(t *testing.T) {
		question := NewQuestion("www.google.com", 1, 1)
		result := question.ToBytes()
		expected := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1}
		if !bytes.Equal(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("should parse question bytes into qname qtype and qclass", func(t *testing.T) {
		data := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1}

		question, _, err := BytesToQuestion(data)
		if err != nil {
			t.Fatalf("BytesToQuestion returned unexpected error: %v", err)
		}

		if question.Qname != "www.google.com" {
			t.Fatalf("Expected Qname %q, got %q", "www.google.com", question.Qname)
		}

		if question.Qtype != 1 {
			t.Fatalf("Expected Qtype 1, got %d", question.Qtype)
		}

		if question.Qclass != 1 {
			t.Fatalf("Expected Qclass 1, got %d", question.Qclass)
		}
	})

	t.Run("should return error for question without null terminator", func(t *testing.T) {
		data := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm'}

		if _, _, err := BytesToQuestion(data); err == nil {
			t.Fatal("Expected error for malformed question without trailing null terminator")
		}
	})
}

func TestReadName(t *testing.T) {
	t.Run("should read an encoded name with labels and terminator", func(t *testing.T) {
		data := []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0}

		name, consumed, err := readName(data, 0)
		if err != nil {
			t.Fatalf("readName returned unexpected error: %v", err)
		}
		if name != "www.google.com" {
			t.Fatalf("Expected name %q, got %q", "www.google.com", name)
		}
		if consumed != len(data) {
			t.Fatalf("Expected consumed %d, got %d", len(data), consumed)
		}
	})

	t.Run("should detect compression pointer and report bytes consumed", func(t *testing.T) {
		data := make([]byte, 32)
		data[0] = 0xc0
		data[1] = 0x0c
		copy(data[12:], []byte{3, 'w', 'w', 'w', 6, 'g', 'o', 'o', 'g', 'l', 'e', 3, 'c', 'o', 'm', 0})

		name, consumed, err := readName(data, 0)
		if err != nil {
			t.Fatalf("readName returned unexpected error: %v", err)
		}
		if name != "www.google.com" {
			t.Fatalf("Expected name %q, got %q", "www.google.com", name)
		}
		if consumed != 2 {
			t.Fatalf("Expected consumed 2, got %d", consumed)
		}
	})

	t.Run("should reject truncated compression pointer", func(t *testing.T) {
		data := []byte{0xc0}

		if _, _, err := readName(data, 0); err == nil {
			t.Fatal("Expected error for truncated compression pointer")
		}
	})

	t.Run("should reject unsupported label encoding", func(t *testing.T) {
		data := []byte{0x80}

		if _, _, err := readName(data, 0); err == nil {
			t.Fatal("Expected error for unsupported label encoding")
		}
	})

	t.Run("should reject label longer than available bytes", func(t *testing.T) {
		data := []byte{0x05, 'a'}

		if _, _, err := readName(data, 0); err == nil {
			t.Fatal("Expected error for label exceeding message length")
		}
	})
}
