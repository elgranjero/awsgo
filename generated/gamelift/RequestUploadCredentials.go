package gamelift

// RequestUploadCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves a fresh set of credentials for use when uploading a new set of game
// build files to Amazon GameLift Servers's Amazon S3. This is done as part of the
// build creation process; see [CreateBuild].
//
// To request new credentials, specify the build ID as returned with an initial
// CreateBuild request. If successful, a new set of credentials are returned, along
// with the S3 storage location associated with the build ID.
//
// # Learn more
//
// [Create a Build with Files in S3]
//
// [All APIs by task]
//
// [Create a Build with Files in S3]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build
// [CreateBuild]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateBuild.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
