package quicksight

// GenerateEmbedUrlForRegisteredUserWithIdentity is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Generates an embed URL that you can use to embed an Amazon Quick Sight
// experience in your website. This action can be used for any type of user that is
// registered in an Amazon Quick Sight account that uses IAM Identity Center for
// authentication. This API requires [identity-enhanced IAM Role sessions]for the authenticated user that the API call
// is being made for.
//
// This API uses [trusted identity propagation] to ensure that an end user is authenticated and receives the
// embed URL that is specific to that user. The IAM Identity Center application
// that the user has logged into needs to have [trusted Identity Propagation enabled for Amazon Quick Sight]with the scope value set to
// quicksight:read . Before you use this action, make sure that you have configured
// the relevant Amazon Quick Sight resource and permissions.
//
// [identity-enhanced IAM Role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-overview.html#types-identity-enhanced-iam-role-sessions
// [trusted Identity Propagation enabled for Amazon Quick Sight]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-using-customermanagedapps-specify-trusted-apps.html
// [trusted identity propagation]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation.html
