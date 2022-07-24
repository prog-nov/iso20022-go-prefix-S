package iso20022

// Choice of format for the repair reason.
type SecuritiesTransactionType35Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType18Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType35Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType18Code)(&value)
}

func (s *SecuritiesTransactionType35Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
