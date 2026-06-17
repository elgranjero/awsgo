package connect

// CreateUser is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates a user account for the specified Amazon Connect instance.
//
// Certain [UserIdentityInfo] parameters are required in some situations. For example, Email ,
// FirstName and LastName are required if you are using Amazon Connect or SAML for
// identity management.
//
// Fields in PhoneConfig cannot be set simultaneously with their corresponding
// channel-specific configuration parameters. Specifically:
//
// - PhoneConfig.AutoAccept conflicts with AutoAcceptConfigs
//
// - PhoneConfig.AfterContactWorkTimeLimit conflicts with AfterContactWorkConfigs
//
// - PhoneConfig.PhoneType and PhoneConfig.PhoneNumber conflict with
// PhoneNumberConfigs
//
// - PhoneConfig.PersistentConnection conflicts with PersistentConnectionConfigs
//
// We recommend using channel-specific parameters such as AutoAcceptConfigs ,
// AfterContactWorkConfigs , PhoneNumberConfigs , PersistentConnectionConfigs , and
// VoiceEnhancementConfigs for per-channel configuration.
//
// For information about how to create users using the Amazon Connect admin
// website, see [Add Users]in the Amazon Connect Administrator Guide.
//
// [Add Users]: https://docs.aws.amazon.com/connect/latest/adminguide/user-management.html
// [UserIdentityInfo]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UserIdentityInfo.html
