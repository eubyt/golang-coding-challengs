package dns

import (
	"bytes"
	"fmt"
)

// Veja https://datatracker.ietf.org/doc/html/rfc1035#section-4.1 para mais detalhes sobre o formato da mensagem DNS.
type DNSMessage struct {
	Header     Header
	Questions  []Question
	Answers    []ResourceRecord
	Authority  []ResourceRecord
	Additional []ResourceRecord
}

func NewDNSMessage(headers Header, questions []Question, records ...[]ResourceRecord) *DNSMessage {
	answers := make([]ResourceRecord, 0)
	authority := make([]ResourceRecord, 0)
	additional := make([]ResourceRecord, 0)

	if len(records) > 0 {
		answers = records[0]
	}
	if len(records) > 1 {
		authority = records[1]
	}
	if len(records) > 2 {
		additional = records[2]
	}

	return &DNSMessage{
		Header:     headers,
		Questions:  questions,
		Answers:    answers,
		Authority:  authority,
		Additional: additional,
	}
}

func BytesToDNSMessage(data []byte) (*DNSMessage, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("data length is less than 12 bytes")
	}

	header, err := BytesToHeader(data[0:12])
	if err != nil {
		return nil, err
	}

	offset := 12

	questions := make([]Question, header.QDCount)
	for i := uint16(0); i < header.QDCount; i++ {
		q, consumed, err := parseQuestionAt(data, offset)
		if err != nil {
			return nil, err
		}
		questions[i] = *q
		offset = consumed
	}

	answers := make([]ResourceRecord, header.ANCount)
	for i := uint16(0); i < header.ANCount; i++ {
		rr, consumed, err := parseResourceRecordAt(data, offset)
		if err != nil {
			return nil, err
		}
		answers[i] = *rr
		offset = consumed
	}

	authority := make([]ResourceRecord, header.NSCount)
	for i := uint16(0); i < header.NSCount; i++ {
		rr, consumed, err := parseResourceRecordAt(data, offset)
		if err != nil {
			return nil, err
		}
		authority[i] = *rr
		offset = consumed
	}

	additional := make([]ResourceRecord, header.ARCount)
	for i := uint16(0); i < header.ARCount; i++ {
		rr, consumed, err := parseResourceRecordAt(data, offset)
		if err != nil {
			return nil, err
		}
		additional[i] = *rr
		offset = consumed
	}

	return &DNSMessage{
		Header:     *header,
		Questions:  questions,
		Answers:    answers,
		Authority:  authority,
		Additional: additional,
	}, nil
}

func (msg *DNSMessage) ToBytes() []byte {
	var buffer bytes.Buffer

	buffer.Write(msg.Header.ToBytes())

	for _, q := range msg.Questions {
		buffer.Write(q.ToBytes())
	}
	for _, a := range msg.Answers {
		buffer.Write(a.ToBytes())
	}
	for _, a := range msg.Authority {
		buffer.Write(a.ToBytes())
	}
	for _, a := range msg.Additional {
		buffer.Write(a.ToBytes())
	}
	return buffer.Bytes()
}

func (msg *DNSMessage) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Header:\n")
	buffer.WriteString(msg.Header.String())
	buffer.WriteString("\n")

	buffer.WriteString("Questions:\n")
	for _, q := range msg.Questions {
		buffer.WriteString(q.String())
		buffer.WriteString("\n")
	}

	buffer.WriteString("Answers:\n")
	for _, a := range msg.Answers {
		buffer.WriteString(a.String())
		buffer.WriteString("\n")
	}

	buffer.WriteString("Authority:\n")
	for _, a := range msg.Authority {
		buffer.WriteString(a.String())
		buffer.WriteString("\n")
	}

	buffer.WriteString("Additional:\n")
	for _, a := range msg.Additional {
		buffer.WriteString(a.String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}
