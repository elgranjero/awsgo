package marketplaceagreement

// SearchAgreements is generated as a reference stub.
// Executable command wiring lives under cmd/marketplaceagreement.go.
//
// Searches across all agreements that a proposer has in AWS Marketplace. The
// search returns a list of agreements with basic agreement information.
//
// The following filter combinations are supported when the PartyType is Proposer :
//
// - AgreementType
//
// - AgreementType + EndTime
//
// - AgreementType + ResourceType
//
// - AgreementType + ResourceType + EndTime
//
// - AgreementType + ResourceType + Status
//
// - AgreementType + ResourceType + Status + EndTime
//
// - AgreementType + ResourceId
//
// - AgreementType + ResourceId + EndTime
//
// - AgreementType + ResourceId + Status
//
// - AgreementType + ResourceId + Status + EndTime
//
// - AgreementType + AcceptorAccountId
//
// - AgreementType + AcceptorAccountId + EndTime
//
// - AgreementType + AcceptorAccountId + Status
//
// - AgreementType + AcceptorAccountId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + OfferId
//
// - AgreementType + AcceptorAccountId + OfferId + Status
//
// - AgreementType + AcceptorAccountId + OfferId + EndTime
//
// - AgreementType + AcceptorAccountId + OfferId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceId
//
// - AgreementType + AcceptorAccountId + ResourceId + Status
//
// - AgreementType + AcceptorAccountId + ResourceId + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceType
//
// - AgreementType + AcceptorAccountId + ResourceType + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceType + Status
//
// - AgreementType + AcceptorAccountId + ResourceType + Status + EndTime
//
// - AgreementType + Status
//
// - AgreementType + Status + EndTime
//
// - AgreementType + OfferId
//
// - AgreementType + OfferId + EndTime
//
// - AgreementType + OfferId + Status
//
// - AgreementType + OfferId + Status + EndTime
//
// - AgreementType + OfferSetId
//
// - AgreementType + OfferSetId + EndTime
//
// - AgreementType + OfferSetId + Status
//
// - AgreementType + OfferSetId + Status + EndTime
//
// To filter by EndTime , you can use either BeforeEndTime or AfterEndTime . Only
// EndTime is supported for sorting.
