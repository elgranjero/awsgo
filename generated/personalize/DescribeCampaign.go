package personalize

// DescribeCampaign is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Describes the given campaign, including its status.
//
// A campaign can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is CREATE FAILED , the response includes the failureReason key,
// which describes why.
//
// For more information on campaigns, see [CreateCampaign].
//
// [CreateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html
