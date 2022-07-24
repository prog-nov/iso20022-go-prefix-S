package iso20022

// Specifies settlement parties (delivering/receiving).
type SettlementParties15 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification47 `xml:"Dpstry"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount51 `xml:"Pty1"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount51 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount51 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties15) AddDepository() *PartyIdentification47 {
	s.Depository = new(PartyIdentification47)
	return s.Depository
}

func (s *SettlementParties15) AddParty1() *PartyIdentificationAndAccount51 {
	s.Party1 = new(PartyIdentificationAndAccount51)
	return s.Party1
}

func (s *SettlementParties15) AddParty2() *PartyIdentificationAndAccount51 {
	s.Party2 = new(PartyIdentificationAndAccount51)
	return s.Party2
}

func (s *SettlementParties15) AddParty3() *PartyIdentificationAndAccount51 {
	s.Party3 = new(PartyIdentificationAndAccount51)
	return s.Party3
}
