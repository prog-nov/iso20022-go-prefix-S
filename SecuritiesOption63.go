package iso20022

// Provides information about the corporate action security option.
type SecuritiesOption63 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate13 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate75 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice64 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties61 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties61 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption63) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption63) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption63) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption63) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption63) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption63) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption63) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption63) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption63) AddPostingQuantity() *Quantity10Choice {
	s.PostingQuantity = new(Quantity10Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption63) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption63) AddFractionDisposition() *FractionDispositionType30Choice {
	s.FractionDisposition = new(FractionDispositionType30Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption63) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption63) AddDateDetails() *SecurityDate13 {
	s.DateDetails = new(SecurityDate13)
	return s.DateDetails
}

func (s *SecuritiesOption63) AddRateDetails() *CorporateActionRate75 {
	s.RateDetails = new(CorporateActionRate75)
	return s.RateDetails
}

func (s *SecuritiesOption63) AddPriceDetails() *CorporateActionPrice64 {
	s.PriceDetails = new(CorporateActionPrice64)
	return s.PriceDetails
}

func (s *SecuritiesOption63) AddReceivingSettlementParties() *SettlementParties61 {
	s.ReceivingSettlementParties = new(SettlementParties61)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption63) AddDeliveringSettlementParties() *SettlementParties61 {
	s.DeliveringSettlementParties = new(SettlementParties61)
	return s.DeliveringSettlementParties
}
