package quicksight

// GetSessionEmbedUrl is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Generates a session URL and authorization code that you can use to embed the
// Amazon Amazon Quick Sight console in your web server code. Use
// GetSessionEmbedUrl where you want to provide an authoring portal that allows
// users to create data sources, datasets, analyses, and dashboards. The users who
// access an embedded Amazon Quick Sight console need belong to the author or admin
// security cohort. If you want to restrict permissions to some of these features,
// add a custom permissions profile to the user with the [UpdateUser]API operation. Use [RegisterUser] API
// operation to add a new user with a custom permission profile attached. For more
// information, see the following sections in the Amazon Quick Suite User Guide:
//
// [Embedding Analytics]
//
// [Customizing Access to the Amazon Quick Suite Console]
//
// [UpdateUser]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateUser.html
// [Customizing Access to the Amazon Quick Suite Console]: https://docs.aws.amazon.com/quicksight/latest/user/customizing-permissions-to-the-quicksight-console.html
// [RegisterUser]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RegisterUser.html
// [Embedding Analytics]: https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html
