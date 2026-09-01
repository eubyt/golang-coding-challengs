package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Header struct {
	ID      uint16 // 16 bits ID
	Flags   uint16 // Flags contendo campos como QR, Opcode, AA, TC, RD, RA, Z e RCODE
	QDCount uint16 // Especifica o número de entradas na seção de perguntas
	ANCount uint16 // Especifica o número de entradas na seção de respostas
	NSCount uint16 // Especifica o número de entradas na seção de autoridade
	ARCount uint16 // Especifica o número de entradas na seção de recursos adicionais
}

func NewHeader(id uint16, flags uint16, qdCount uint16, anCount uint16, nsCount uint16, arCount uint16) *Header {
	return &Header{
		ID:      id,
		Flags:   flags,
		QDCount: qdCount,
		ANCount: anCount,
		NSCount: nsCount,
		ARCount: arCount,
	}
}

func BytesToHeader(data []byte) (*Header, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("data length is less than 12 bytes")
	}

	header := &Header{}
	buffer := bytes.NewReader(data)

	binary.Read(buffer, binary.BigEndian, &header.ID)
	binary.Read(buffer, binary.BigEndian, &header.Flags)
	binary.Read(buffer, binary.BigEndian, &header.QDCount)
	binary.Read(buffer, binary.BigEndian, &header.ANCount)
	binary.Read(buffer, binary.BigEndian, &header.NSCount)
	binary.Read(buffer, binary.BigEndian, &header.ARCount)

	return header, nil
}

func (h *Header) ToBytes() []byte {
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, h.ID)
	binary.Write(&buffer, binary.BigEndian, h.Flags)
	binary.Write(&buffer, binary.BigEndian, h.QDCount)
	binary.Write(&buffer, binary.BigEndian, h.ANCount)
	binary.Write(&buffer, binary.BigEndian, h.NSCount)
	binary.Write(&buffer, binary.BigEndian, h.ARCount)

	return buffer.Bytes()
}

func (h *Header) getFlags() *HeaderFlag {
	return Uint16ToHeaderFlag(h.Flags)
}

func (h *Header) String() string {
	flags := h.getFlags()
	return fmt.Sprintf(
		"  ID: %d\n  Flags: %s\n  QDCount: %d\n  ANCount: %d\n  NSCount: %d\n  ARCount: %d",
		h.ID,
		flags.String(),
		h.QDCount, h.ANCount, h.NSCount, h.ARCount,
	)
}
