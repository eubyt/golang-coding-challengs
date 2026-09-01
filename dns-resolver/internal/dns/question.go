package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type Question struct {
	Qname  string // Nome do domínio da pergunta no formato RFC 1035
	Qtype  uint16 // Tipo de registro da pergunta
	Qclass uint16 // Classe do registro da pergunta
}

func NewQuestion(qname string, qtype uint16, qclass uint16) *Question {
	return &Question{
		Qname:  encodeName(qname),
		Qtype:  qtype,
		Qclass: qclass,
	}
}

func BytesToQuestion(data []byte) (*Question, int, error) {
	return parseQuestionAt(data, 0)
}

func parseQuestionAt(data []byte, offset int) (*Question, int, error) {
	if len(data)-offset < 4 {
		return nil, 0, fmt.Errorf("data length is less than 4 bytes")
	}

	name, consumedRelative, err := readName(data, offset)
	if err != nil {
		return nil, 0, err
	}
	consumed := offset + consumedRelative
	if len(data)-consumed < 4 {
		return nil, 0, fmt.Errorf("invalid question format: missing QTYPE/QCLASS")
	}

	question := &Question{Qname: name}
	question.Qtype = binary.BigEndian.Uint16(data[consumed : consumed+2])
	question.Qclass = binary.BigEndian.Uint16(data[consumed+2 : consumed+4])
	return question, consumed + 4, nil
}

func (q *Question) ToBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(q.Qname)
	binary.Write(&buffer, binary.BigEndian, q.Qtype)
	binary.Write(&buffer, binary.BigEndian, q.Qclass)

	return buffer.Bytes()
}

func (q *Question) String() string {
	return fmt.Sprintf(
		"  Qname: %s, Qtype: %s, Qclass: %s",
		q.Qname,
		DNSRecordType(q.Qtype).String(),
		DNSRecordClass(q.Qclass).String(),
	)
}

func readName(data []byte, offset int) (string, int, error) {
	if offset >= len(data) {
		return "", 0, fmt.Errorf("invalid domain name: offset out of range")
	}

	start := offset
	labels := make([]string, 0, 4)
	for {
		if offset >= len(data) {
			return "", 0, fmt.Errorf("invalid domain name: missing terminator")
		}

		length := int(data[offset])
		if length == 0 {
			if len(labels) == 0 {
				return "", offset + 1 - start, nil
			}
			return strings.Join(labels, "."), offset + 1 - start, nil
		}
		if length&0xC0 == 0xC0 {
			if offset+1 >= len(data) {
				return "", 0, fmt.Errorf("invalid domain name: truncated compression pointer")
			}

			ptr := int(binary.BigEndian.Uint16(data[offset:offset+2])) & 0x3FFF
			suffix, _, err := readName(data, ptr)
			if err != nil {
				return "", 0, err
			}
			if len(labels) == 0 {
				return suffix, offset + 2 - start, nil
			}
			return strings.Join(append(labels, suffix), "."), offset + 2 - start, nil
		}
		if length&0xC0 != 0 {
			return "", 0, fmt.Errorf("invalid domain name: unsupported label encoding")
		}
		if offset+1+length > len(data) {
			return "", 0, fmt.Errorf("invalid domain name: label exceeds message length")
		}

		labels = append(labels, string(data[offset+1:offset+1+length]))
		offset += 1 + length
	}
}

// encodeName converte o nome do domínio para o formato de
// {tamanho + nome do rótulo} conforme especificado na RFC 1035.
func encodeName(name string) string {
	var encoded []byte
	domainParts := strings.Split(name, ".")
	for _, part := range domainParts {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, []byte(part)...)
	}
	encoded = append(encoded, 0) // Adiciona o byte nulo no final
	return string(encoded)
}
