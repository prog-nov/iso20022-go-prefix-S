package iso20022

// Tax related to an investment fund order.
type Tax21 struct {

	// Type of tax.
	Type *TaxType1Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, for example, the purchase price.
	Basis *TaxBasis1Choice `xml:"Bsis"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation8 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax21) AddType() *TaxType1Choice {
	t.Type = new(TaxType1Choice)
	return t.Type
}

func (t *Tax21) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax21) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *Tax21) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax21) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax21) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax21) AddTaxCalculationDetails() *TaxCalculationInformation8 {
	t.TaxCalculationDetails = new(TaxCalculationInformation8)
	return t.TaxCalculationDetails
}
