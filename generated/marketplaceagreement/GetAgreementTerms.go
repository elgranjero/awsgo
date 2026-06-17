package marketplaceagreement

// GetAgreementTerms is generated as a reference stub.
// Executable command wiring lives under cmd/marketplaceagreement.go.
//
// Obtains details about the terms in an agreement that you participated in as
// proposer or acceptor.
//
// The details include:
//
// - TermType – The type of term, such as LegalTerm , RenewalTerm , or
// ConfigurableUpfrontPricingTerm .
//
// - TermID – The ID of the particular term, which is common between offer and
// agreement.
//
// - TermPayload – The key information contained in the term, such as the EULA
// for LegalTerm or pricing and dimensions for various pricing terms, such as
// ConfigurableUpfrontPricingTerm or UsageBasedPricingTerm .
//
// - Configuration – The buyer/acceptor's selection at the time of agreement
// creation, such as the number of units purchased for a dimension or setting the
// EnableAutoRenew flag.
