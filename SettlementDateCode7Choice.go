package iso20022

// Choice of format for the settlement date code.
type SettlementDateCode7Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *SettlementDate4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementDateCode7Choice) SetCode(value string) {
	s.Code = (*SettlementDate4Code)(&value)
}

func (s *SettlementDateCode7Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
