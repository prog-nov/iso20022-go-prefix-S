package iso20022

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation5 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation5 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation5) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation5) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation5) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation5) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation5) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
