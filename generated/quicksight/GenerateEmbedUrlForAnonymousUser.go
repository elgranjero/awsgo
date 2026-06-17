package quicksight

// GenerateEmbedUrlForAnonymousUser is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Generates an embed URL that you can use to embed an Amazon Quick Suite
// dashboard or visual in your website, without having to register any reader
// users. Before you use this action, make sure that you have configured the
// dashboards and permissions.
//
// The following rules apply to the generated URL:
//
// - It contains a temporary bearer token. It is valid for 5 minutes after it is
// generated. Once redeemed within this period, it cannot be re-used again.
//
// - The URL validity period should not be confused with the actual session
// lifetime that can be customized using the [SessionLifetimeInMinutes]parameter. The resulting user
// session is valid for 15 minutes (minimum) to 10 hours (maximum). The default
// session duration is 10 hours.
//
// - You are charged only when the URL is used or there is interaction with
// Amazon Quick Suite.
//
// For more information, see [Embedded Analytics] in the Amazon Quick Suite User Guide.
//
// For more information about the high-level steps for embedding and for an
// interactive demo of the ways you can customize embedding, visit the [Amazon Quick Suite Developer Portal].
//
// [Embedded Analytics]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html
// [Amazon Quick Suite Developer Portal]: https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html
// [SessionLifetimeInMinutes]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GenerateEmbedUrlForAnonymousUser.html#QS-GenerateEmbedUrlForAnonymousUser-request-SessionLifetimeInMinutes
