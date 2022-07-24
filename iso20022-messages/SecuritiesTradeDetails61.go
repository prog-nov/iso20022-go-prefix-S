package iso20022

// Details of the securities trade.
type SecuritiesTradeDetails61 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus4Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount50 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails34 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails108 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection71 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts36 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties31 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails61) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails61) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails61) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails61) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails61) AddStatus() *AllegementStatus4Choice {
	s.Status = new(AllegementStatus4Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails61) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails61) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails61) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails61) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails61) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails61) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails61) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails61) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails61) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails61) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails61) AddQuantityAndAccountDetails() *QuantityAndAccount50 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount50)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails61) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails34 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails34)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails61) AddSettlementParameters() *SettlementDetails108 {
	s.SettlementParameters = new(SettlementDetails108)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails61) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails61) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails61) AddSettlementAmount() *AmountAndDirection71 {
	s.SettlementAmount = new(AmountAndDirection71)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails61) AddOtherAmounts() *OtherAmounts36 {
	s.OtherAmounts = new(OtherAmounts36)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails61) AddOtherBusinessParties() *OtherParties31 {
	s.OtherBusinessParties = new(OtherParties31)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails61) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
