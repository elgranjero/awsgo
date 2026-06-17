package personalize

// CreateCampaign is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// You incur campaign costs while it is active. To avoid unnecessary costs, make
//
// sure to delete the campaign when you are finished. For information about
// campaign costs, see [Amazon Personalize pricing].
//
// Creates a campaign that deploys a solution version. When a client calls the [GetRecommendations]
// and [GetPersonalizedRanking]APIs, a campaign is specified in the request.
//
// # Minimum Provisioned TPS and Auto-Scaling
//
// A high minProvisionedTPS will increase your cost. We recommend starting with 1
// for minProvisionedTPS (the default). Track your usage using Amazon CloudWatch
// metrics, and increase the minProvisionedTPS as necessary.
//
// When you create an Amazon Personalize campaign, you can specify the minimum
// provisioned transactions per second ( minProvisionedTPS ) for the campaign. This
// is the baseline transaction throughput for the campaign provisioned by Amazon
// Personalize. It sets the minimum billing charge for the campaign while it is
// active. A transaction is a single GetRecommendations or GetPersonalizedRanking
// request. The default minProvisionedTPS is 1.
//
// If your TPS increases beyond the minProvisionedTPS , Amazon Personalize
// auto-scales the provisioned capacity up and down, but never below
// minProvisionedTPS . There's a short time delay while the capacity is increased
// that might cause loss of transactions. When your traffic reduces, capacity
// returns to the minProvisionedTPS .
//
// You are charged for the the minimum provisioned TPS or, if your requests exceed
// the minProvisionedTPS , the actual TPS. The actual TPS is the total number of
// recommendation requests you make. We recommend starting with a low
// minProvisionedTPS , track your usage using Amazon CloudWatch metrics, and then
// increase the minProvisionedTPS as necessary.
//
// For more information about campaign costs, see [Amazon Personalize pricing].
//
// # Status
//
// A campaign can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the campaign status, call [DescribeCampaign].
//
// Wait until the status of the campaign is ACTIVE before asking the campaign for
// recommendations.
//
// # Related APIs
//
// [ListCampaigns]
//
// [DescribeCampaign]
//
// [UpdateCampaign]
//
// [DeleteCampaign]
//
// [UpdateCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateCampaign.html
// [GetRecommendations]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html
// [ListCampaigns]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html
// [DeleteCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteCampaign.html
// [GetPersonalizedRanking]: https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetPersonalizedRanking.html
// [Amazon Personalize pricing]: https://aws.amazon.com/personalize/pricing/
// [DescribeCampaign]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeCampaign.html
