package iso20022

// Information related to an identification, eg, party identification or account identification.
type SimpleIdentificationInformation struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}
