package quicksight

// PredictQAResults is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Predicts existing visuals or generates new visuals to answer a given query.
//
// This API uses [trusted identity propagation] to ensure that an end user is authenticated and receives the
// embed URL that is specific to that user. The IAM Identity Center application
// that the user has logged into needs to have [trusted Identity Propagation enabled for Quick Suite]with the scope value set to
// quicksight:read . Before you use this action, make sure that you have configured
// the relevant Quick Suite resource and permissions.
//
// We recommend enabling the QSearchStatus API to unlock the full potential of
// PredictQnA . When QSearchStatus is enabled, it first checks the specified
// dashboard for any existing visuals that match the question. If no matching
// visuals are found, PredictQnA uses generative Q&A to provide an answer. To
// update the QSearchStatus , see [UpdateQuickSightQSearchConfiguration].
//
// [trusted Identity Propagation enabled for Quick Suite]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-using-customermanagedapps-specify-trusted-apps.html
// [trusted identity propagation]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation.html
// [UpdateQuickSightQSearchConfiguration]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateQuickSightQSearchConfiguration.html
