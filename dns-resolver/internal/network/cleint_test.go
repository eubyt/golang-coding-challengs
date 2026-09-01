package network

import "testing"

func TestClient_getTypeIp(t *testing.T) {
	t.Run("should return IPv4 for valid IPv4 address", func(t *testing.T) {
		client := NewClient("8.8.8.8", 53)
		ipType, err := client.getTypeIp()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if ipType != "IPv4" {
			t.Errorf("Expected IPv4, but got %s", ipType)
		}
	})

	t.Run("should return IPv6 for valid IPv6 address", func(t *testing.T) {
		client := NewClient("2001:4860:4860::8888", 53)
		ipType, err := client.getTypeIp()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if ipType != "IPv6" {
			t.Errorf("Expected IPv6, but got %s", ipType)
		}
	})

	t.Run("should return error for invalid IP address", func(t *testing.T) {
		client := NewClient("invalid_ip", 53)
		_, err := client.getTypeIp()
		if err == nil {
			t.Errorf("Expected error for invalid IP address, but got nil")
		}
	})
}

func TestClient_parseAddress(t *testing.T) {
	t.Run("should format IPv4 address correctly", func(t *testing.T) {
		client := NewClient("8.8.8.8", 53)
		address, err := client.parseAddress()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := "8.8.8.8:53"
		if address != expected {
			t.Errorf("Expected %s, but got %s", expected, address)
		}
	})

	t.Run("should format IPv6 address correctly", func(t *testing.T) {
		client := NewClient("2001:4860:4860::8888", 53)
		address, err := client.parseAddress()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := "[2001:4860:4860::8888]:53"
		if address != expected {
			t.Errorf("Expected %s, but got %s", expected, address)
		}
	})

	t.Run("should return error for invalid IP address", func(t *testing.T) {
		client := NewClient("invalid_ip", 53)
		_, err := client.parseAddress()
		if err == nil {
			t.Errorf("Expected error for invalid IP address, but got nil")
		}
	})
}
