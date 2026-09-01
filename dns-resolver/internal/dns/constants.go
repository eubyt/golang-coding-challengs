package dns

type DNSQRCode bool

const (
	TypeResponse DNSQRCode = true
	TypeQuery    DNSQRCode = false
)

func (q DNSQRCode) String() string {
	if q {
		return "response"
	}
	return "query"
}

type DNSOpcode uint8

const (
	OpcodeQuery DNSOpcode = iota
	OpcodeIQuery
	OpcodeStatus
)

func (o DNSOpcode) String() string {
	switch o {
	case OpcodeQuery:
		return "QUERY"
	case OpcodeIQuery:
		return "IQUERY"
	case OpcodeStatus:
		return "STATUS"
	default:
		return "UNKNOWN"
	}
}

type DNSRCode uint8

const (
	RCodeNoError DNSRCode = iota
	RCodeFormErr
	RCodeServFail
	RCodeNXDomain
	RCodeNotImp
	RCodeRefused
)

func (r DNSRCode) String() string {
	switch r {
	case RCodeNoError:
		return "NOERROR"
	case RCodeFormErr:
		return "FORMERR"
	case RCodeServFail:
		return "SERVFAIL"
	case RCodeNXDomain:
		return "NXDOMAIN"
	case RCodeNotImp:
		return "NOTIMP"
	case RCodeRefused:
		return "REFUSED"
	default:
		return "UNKNOWN"
	}
}

// DNS record types
type DNSRecordType uint16

const (
	TypeA     DNSRecordType = 1   // IPv4 address record
	TypeNS    DNSRecordType = 2   // authoritative name server record
	TypeCNAME DNSRecordType = 5   // canonical name record
	TypeSOA   DNSRecordType = 6   // start of authority record
	TypePTR   DNSRecordType = 12  // pointer record
	TypeMX    DNSRecordType = 15  // mail exchange record
	TypeTXT   DNSRecordType = 16  // text record
	TypeAAAA  DNSRecordType = 28  // IPv6 address record
	TypeSRV   DNSRecordType = 33  // service locator record
	TypeOPT   DNSRecordType = 41  // option record
	TypeAXFR  DNSRecordType = 252 // transfer of an entire zone record
	TypeMAILB DNSRecordType = 253 // mailbox-related records (MB, MG, MR)
	TypeMAILA DNSRecordType = 254 // mail agent RRs (Obsolete - see MX)
	TypeAll   DNSRecordType = 255 // all records
)

func (rt DNSRecordType) String() string {
	switch rt {
	case TypeA:
		return "A"
	case TypeNS:
		return "NS"
	case TypeCNAME:
		return "CNAME"
	case TypeSOA:
		return "SOA"
	case TypePTR:
		return "PTR"
	case TypeMX:
		return "MX"
	case TypeTXT:
		return "TXT"
	case TypeAAAA:
		return "AAAA"
	case TypeSRV:
		return "SRV"
	case TypeOPT:
		return "OPT"
	case TypeAXFR:
		return "AXFR"
	case TypeMAILB:
		return "MAILB"
	case TypeMAILA:
		return "MAILA"
	case TypeAll:
		return "ALL"
	default:
		return "UNKNOWN"
	}
}

// DNS record classes
type DNSRecordClass uint16

const (
	ClassIN  DNSRecordClass = 1   // Internet class
	ClassCS  DNSRecordClass = 2   // CSNET class (Obsolete)
	ClassCH  DNSRecordClass = 3   // CHAOS class
	ClassHS  DNSRecordClass = 4   // Hesiod [Dyer 87]
	ClassAll DNSRecordClass = 255 // all classes
)

func (rc DNSRecordClass) String() string {
	switch rc {
	case ClassIN:
		return "IN"
	case ClassCS:
		return "CS"
	case ClassCH:
		return "CH"
	case ClassHS:
		return "HS"
	case ClassAll:
		return "ALL"
	default:
		return "UNKNOWN"
	}
}
