package iso20022

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters12 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters12) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters12) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters12) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}
