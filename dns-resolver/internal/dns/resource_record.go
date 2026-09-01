package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ResourceRecord struct {
	Name     string // Nome do recurso
	Type     uint16 // Tipo do recurso (A, AAAA, CNAME, etc.)
	Class    uint16 // Classe do recurso (IN, CH, etc.)
	TTL      uint32 // Tempo de vida do recurso em segundos
	RDLength uint16 // Comprimento dos dados do recurso
	RData    []byte // Dados do recurso (endereço IP, nome canônico, etc.)
}

func NewResourceRecord(name string, recordType DNSRecordType, class DNSRecordClass, ttl uint32, rdLength uint16, rdata []byte) *ResourceRecord {
	return &ResourceRecord{
		Name:     encodeName(name),
		Type:     uint16(recordType),
		Class:    uint16(class),
		TTL:      ttl,
		RDLength: rdLength,
		RData:    rdata,
	}
}

func BytesToResourceRecord(data []byte) (*ResourceRecord, int, error) {
	return parseResourceRecordAt(data, 0)
}

func parseResourceRecordAt(data []byte, offset int) (*ResourceRecord, int, error) {
	if len(data)-offset < 10 {
		return nil, 0, fmt.Errorf("data length is less than 10 bytes")
	}

	name, consumedRelative, err := readName(data, offset)
	if err != nil {
		return nil, 0, err
	}
	base := offset + consumedRelative
	if len(data)-base < 10 {
		return nil, 0, fmt.Errorf("invalid resource record format: missing record metadata")
	}

	rr := &ResourceRecord{Name: name}
	rr.Type = binary.BigEndian.Uint16(data[base : base+2])
	rr.Class = binary.BigEndian.Uint16(data[base+2 : base+4])
	rr.TTL = binary.BigEndian.Uint32(data[base+4 : base+8])
	rr.RDLength = binary.BigEndian.Uint16(data[base+8 : base+10])

	if len(data)-base < 10+int(rr.RDLength) {
		return nil, 0, fmt.Errorf("invalid resource record format: RData exceeds message length")
	}

	rr.RData = make([]byte, rr.RDLength)
	copy(rr.RData, data[base+10:base+10+int(rr.RDLength)])
	return rr, base + 10 + int(rr.RDLength), nil
}

func (rr *ResourceRecord) ToBytes() []byte {
	var buffer bytes.Buffer

	buffer.WriteString(rr.Name)
	binary.Write(&buffer, binary.BigEndian, rr.Type)
	binary.Write(&buffer, binary.BigEndian, rr.Class)
	binary.Write(&buffer, binary.BigEndian, rr.TTL)
	binary.Write(&buffer, binary.BigEndian, rr.RDLength)
	buffer.Write(rr.RData)

	return buffer.Bytes()
}

func (rr *ResourceRecord) String() string {
	return fmt.Sprintf(
		"  Name: %s, Type: %s, Class: %s, TTL: %d, RDLength: %d, RData: %v",
		rr.Name,
		DNSRecordType(rr.Type).String(),
		DNSRecordClass(rr.Class).String(),
		rr.TTL,
		rr.RDLength,
		rr.RData,
	)
}
