package quicksight

// GetDashboardEmbedUrl is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Generates a temporary session URL and authorization code(bearer token) that you
// can use to embed an Amazon Quick Sight read-only dashboard in your website or
// application. Before you use this command, make sure that you have configured the
// dashboards and permissions.
//
// Currently, you can use GetDashboardEmbedURL only from the server, not from the
// user's browser. The following rules apply to the generated URL:
//
// - They must be used together.
//
// - They can be used one time only.
//
// - They are valid for 5 minutes after you run this command.
//
// - You are charged only when the URL is used or there is interaction with
// Quick Suite.
//
// - The resulting user session is valid for 15 minutes (default) up to 10 hours
// (maximum). You can use the optional SessionLifetimeInMinutes parameter to
// customize session duration.
//
// For more information, see [Embedding Analytics Using GetDashboardEmbedUrl] in the Amazon Quick Suite User Guide.
//
// For more information about the high-level steps for embedding and for an
// interactive demo of the ways you can customize embedding, visit the [Amazon Quick Suite Developer Portal].
//
// [Amazon Quick Suite Developer Portal]: https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html
// [Embedding Analytics Using GetDashboardEmbedUrl]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics-deprecated.html
