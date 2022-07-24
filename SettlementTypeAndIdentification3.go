package iso20022

// Provides transaction type and identification information.
type SettlementTypeAndIdentification3 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification3) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification3) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification3) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}
