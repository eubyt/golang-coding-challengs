package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eubyt/codingchallenges/dns-resolver/internal/dns"
	"github.com/eubyt/codingchallenges/dns-resolver/internal/network"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run ./cmd <domain>")
		return
	}

	domain := strings.TrimSpace(os.Args[1])
	if domain == "" {
		fmt.Println("Domain cannot be empty")
		return
	}

	id := uint16(12345)
	flags := dns.NewHeaderFlag(dns.TypeQuery, 0, false, false, true, false, 0, dns.RCodeNoError).ToUint16()
	header := dns.NewHeader(id, flags, 1, 0, 0, 0)

	question := dns.NewQuestion(domain, 1, 1)
	msg := dns.NewDNSMessage(*header, []dns.Question{*question})

	client := network.NewClient("8.8.8.8", 53)

	resp, err := client.Query(msg.ToBytes())
	if err != nil {
		fmt.Println("query error:", err)
		return
	}

	dnsMsg, err := dns.BytesToDNSMessage(resp)
	if err != nil {
		fmt.Println("error parsing DNS message:", err)
		return
	}

	fmt.Print(dnsMsg.String())
}
