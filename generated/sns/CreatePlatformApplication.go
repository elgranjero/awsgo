package sns

// CreatePlatformApplication is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Creates a platform application object for one of the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging), to which
// devices and mobile apps may register. You must specify PlatformPrincipal and
// PlatformCredential attributes when using the CreatePlatformApplication action.
//
// PlatformPrincipal and PlatformCredential are received from the notification
// service.
//
// - For ADM, PlatformPrincipal is client id and PlatformCredential is client
// secret .
//
// - For APNS and APNS_SANDBOX using certificate credentials, PlatformPrincipal
// is SSL certificate and PlatformCredential is private key .
//
// - For APNS and APNS_SANDBOX using token credentials, PlatformPrincipal is
// signing key ID and PlatformCredential is signing key .
//
// - For Baidu, PlatformPrincipal is API key and PlatformCredential is secret key
// .
//
// - For GCM (Firebase Cloud Messaging) using key credentials, there is no
// PlatformPrincipal . The PlatformCredential is API key .
//
// - For GCM (Firebase Cloud Messaging) using token credentials, there is no
// PlatformPrincipal . The PlatformCredential is a JSON formatted private key
// file. When using the Amazon Web Services CLI or Amazon Web Services SDKs, the
// file must be in string format and special characters must be ignored. To format
// the file correctly, Amazon SNS recommends using the following command:
// SERVICE_JSON=$(jq (at)json < service.json) .
//
// - For MPNS, PlatformPrincipal is TLS certificate and PlatformCredential is
// private key .
//
// - For WNS, PlatformPrincipal is Package Security Identifier and
// PlatformCredential is secret key .
//
// You can use the returned PlatformApplicationArn as an attribute for the
// CreatePlatformEndpoint action.
