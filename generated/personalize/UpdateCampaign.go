package personalize

// UpdateCampaign is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Updates a campaign to deploy a retrained solution version with an existing
//
// campaign, change your campaign's minProvisionedTPS , or modify your campaign's
// configuration. For example, you can set enableMetadataWithRecommendations to
// true for an existing campaign.
//
// To update a campaign to start automatically using the latest solution version,
// specify the following:
//
// - For the SolutionVersionArn parameter, specify the Amazon Resource Name (ARN)
// of your solution in SolutionArn/$LATEST format.
//
// - In the campaignConfig , set syncWithLatestSolutionVersion to true .
//
// To update a campaign, the campaign status must be ACTIVE or CREATE FAILED.
// Check the campaign status using the [DescribeCampaign]operation.
//
// You can still get recommendations from a campaign while an update is in
// progress. The campaign will use the previous solution version and campaign
// configuration to generate recommendations until the latest campaign update
// status is Active .
//
// For more information about updating a campaign, including code samples, see [Updating a campaign].
// For more information about campaigns, see [Creating a campaign].
//
// [Creating a campaign]: https://docs.aws.amazon.com/personalize/latest/dg/campaigns.html
// [Updating a campaign]: https://docs.aws.amazon.com/personalize/latest/dg/update-campaigns.html
// [DescribeCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeCampaign.html
