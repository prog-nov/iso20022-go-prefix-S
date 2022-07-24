package iso20022

// Specifies security date details.
type SecurityDate5 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat19Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat19Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat19Choice `xml:"PrpssDt,omitempty"`
}

func (s *SecurityDate5) AddPaymentDate() *DateFormat19Choice {
	s.PaymentDate = new(DateFormat19Choice)
	return s.PaymentDate
}

func (s *SecurityDate5) AddAvailableDate() *DateFormat19Choice {
	s.AvailableDate = new(DateFormat19Choice)
	return s.AvailableDate
}

func (s *SecurityDate5) AddDividendRankingDate() *DateFormat19Choice {
	s.DividendRankingDate = new(DateFormat19Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate5) AddEarliestPaymentDate() *DateFormat19Choice {
	s.EarliestPaymentDate = new(DateFormat19Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate5) AddPariPassuDate() *DateFormat19Choice {
	s.PariPassuDate = new(DateFormat19Choice)
	return s.PariPassuDate
}
