package dns

import (
	"fmt"

	"github.com/eubyt/codingchallenges/dns-resolver/internal/util"
)

type HeaderFlag struct {
	QR     bool  // true = resposta; false = consulta.
	Opcode uint8 // Código da operação DNS (OPCODE): 0 = QUERY, 1 = IQUERY, 2 = STATUS; 3-15 reservados para uso futuro.
	AA     bool  // true = servidor que responde é uma autoridade para o domínio consultado; false = nao é autoridade.
	TC     bool  // true = mensagem truncada; false = mensagem nao truncada.
	RD     bool  // true = cliente solicita recursao; false = nao solicita recursao.
	RA     bool  // true = servidor oferece suporte a consultas recursivas; false = nao oferece suporte.
	Z      uint8 // Campo reservado para uso futuro; deve permanecer igual a zero.
	RCODE  uint8 // Código de resposta DNS (RCODE): 0 = NOERROR, 1 = FORMERR, 2 = SERVFAIL, 3 = NXDOMAIN, 4 = NOTIMP, 5 = REFUSED.
}

func NewHeaderFlag(isResponse DNSQRCode, opcode DNSOpcode, isAuthoritative bool, isTruncated bool, recursionDesired bool, recursionAvailable bool, reservedZone uint8, responseCode DNSRCode) *HeaderFlag {
	return &HeaderFlag{
		QR:     bool(isResponse),
		Opcode: uint8(opcode),
		AA:     isAuthoritative,
		TC:     isTruncated,
		RD:     recursionDesired,
		RA:     recursionAvailable,
		Z:      reservedZone,
		RCODE:  uint8(responseCode),
	}
}

func Uint16ToHeaderFlag(value uint16) *HeaderFlag {
	qr := (value >> 15) & 1
	opcode := (value >> 11) & 0x0F
	aa := (value >> 10) & 1
	tc := (value >> 9) & 1
	rd := (value >> 8) & 1
	ra := (value >> 7) & 1
	z := (value >> 4) & 0x07
	rcode := value & 0x0F

	return &HeaderFlag{
		QR:     qr == 1,
		Opcode: uint8(opcode),
		AA:     aa == 1,
		TC:     tc == 1,
		RD:     rd == 1,
		RA:     ra == 1,
		Z:      uint8(z),
		RCODE:  uint8(rcode),
	}
}

func (hf *HeaderFlag) ToUint16() uint16 {
	qr := uint16(util.BoolToInt(hf.QR)) << 15 // 32768 OU 0
	opcode := uint16(hf.Opcode) << 11         // 2048, 4096, 6144 OU 0
	aa := uint16(util.BoolToInt(hf.AA)) << 10 // 1024 OU 0
	tc := uint16(util.BoolToInt(hf.TC)) << 9  // 512 OU 0
	rd := uint16(util.BoolToInt(hf.RD)) << 8  // 256 OU 0
	ra := uint16(util.BoolToInt(hf.RA)) << 7  // 128 OU 0
	z := uint16(hf.Z) << 4
	rcode := uint16(hf.RCODE)

	return uint16(qr | opcode | aa | tc | rd | ra | z | rcode)
}

func (hf *HeaderFlag) qrString() string {
	return DNSQRCode(hf.QR).String()
}

func (hf *HeaderFlag) String() string {
	return fmt.Sprintf(
		"QR=%s, Opcode=%s, AA=%t, TC=%t, RD=%t, RA=%t, Z=%d, RCODE=%s",
		hf.qrString(),
		DNSOpcode(hf.Opcode).String(),
		hf.AA,
		hf.TC,
		hf.RD,
		hf.RA,
		hf.Z,
		DNSRCode(hf.RCODE).String(),
	)
}
