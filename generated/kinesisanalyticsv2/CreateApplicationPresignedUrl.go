package kinesisanalyticsv2

// CreateApplicationPresignedUrl is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Creates and returns a URL that you can use to connect to an application's
// extension.
//
// The IAM role or user used to call this API defines the permissions to access
// the extension. After the presigned URL is created, no additional permission is
// required to access this URL. IAM authorization policies for this API are also
// enforced for every HTTP request that attempts to connect to the extension.
//
// You control the amount of time that the URL will be valid using the
// SessionExpirationDurationInSeconds parameter. If you do not provide this
// parameter, the returned URL is valid for twelve hours.
//
// The URL that you get from a call to CreateApplicationPresignedUrl must be used
// within 3 minutes to be valid. If you first try to use the URL after the 3-minute
// limit expires, the service returns an HTTP 403 Forbidden error.
