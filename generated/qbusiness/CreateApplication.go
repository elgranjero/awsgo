package qbusiness

// CreateApplication is generated as a reference stub.
// Executable command wiring lives under cmd/qbusiness.go.
//
// Creates an Amazon Q Business application.
//
// There are new tiers for Amazon Q Business. Not all features in Amazon Q
// Business Pro are also available in Amazon Q Business Lite. For information on
// what's included in Amazon Q Business Lite and what's included in Amazon Q
// Business Pro, see [Amazon Q Business tiers]. You must use the Amazon Q Business console to assign
// subscription tiers to users.
//
// An Amazon Q Apps service linked role will be created if it's absent in the
// Amazon Web Services account when QAppsConfiguration is enabled in the request.
// For more information, see [Using service-linked roles for Q Apps].
//
// When you create an application, Amazon Q Business may securely transmit data
// for processing from your selected Amazon Web Services region, but within your
// geography. For more information, see [Cross region inference in Amazon Q Business].
//
// [Amazon Q Business tiers]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#user-sub-tiers
// [Using service-linked roles for Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
