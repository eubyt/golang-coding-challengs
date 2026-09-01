package network

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	ipAddress string
	port      int
	timeout   time.Duration
}

func NewClient(addr string, port int) *Client {
	return &Client{
		ipAddress: addr,
		port:      port,
		timeout:   5 * time.Second,
	}
}

// getTypeIp para determinar se o endereço IP é IPv4 ou IPv6
func (c *Client) getTypeIp() (string, error) {
	ip := net.ParseIP(c.ipAddress)
	if ip.To4() != nil {
		return "IPv4", nil
	} else if ip.To16() != nil {
		return "IPv6", nil
	}

	return "", fmt.Errorf("invalid IP address: %s", c.ipAddress)
}

// parseAddress para formatar o address de acordo com o tipo retornado por getTypeIp
func (c *Client) parseAddress() (string, error) {
	ipType, err := c.getTypeIp()
	if err != nil {
		return "", err
	}

	switch ipType {
	case "IPv4":
		return fmt.Sprintf("%s:%d", c.ipAddress, c.port), nil
	case "IPv6":
		return fmt.Sprintf("[%s]:%d", c.ipAddress, c.port), nil
	default:
		return "", fmt.Errorf("unsupported IP type: %s", ipType)
	}
}

func (c *Client) Query(message []byte) ([]byte, error) {
	address, err := c.parseAddress()
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", address, err)
	}

	defer conn.Close()

	conn.SetDeadline(time.Now().Add(c.timeout))

	_, err = conn.Write(message)
	if err != nil {
		return nil, fmt.Errorf("failed to send message to %s: %v", address, err)
	}

	// Read the response from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from %s: %v", address, err)
	}

	response := buffer[:n]

	return response, nil
}
