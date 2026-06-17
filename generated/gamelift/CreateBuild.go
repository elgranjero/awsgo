package gamelift

// CreateBuild is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere
//
// Creates a new Amazon GameLift Servers build resource for your game server
// binary files. Combine game server binaries into a zip file for use with Amazon
// GameLift Servers.
//
// When setting up a new game build for Amazon GameLift Servers, we recommend
// using the CLI command [upload-build]. This helper command combines two tasks: (1) it uploads
// your build files from a file directory to an Amazon GameLift Servers Amazon S3
// location, and (2) it creates a new build resource.
//
// You can use the CreateBuild operation in the following scenarios:
//
// - Create a new game build with build files that are in an Amazon S3 location
// under an Amazon Web Services account that you control. To use this option, you
// give Amazon GameLift Servers access to the Amazon S3 bucket. With permissions in
// place, specify a build name, operating system, and the Amazon S3 storage
// location of your game build.
//
// - Upload your build files to a Amazon GameLift Servers Amazon S3 location. To
// use this option, specify a build name and operating system. This operation
// creates a new build resource and also returns an Amazon S3 location with
// temporary access credentials. Use the credentials to manually upload your build
// files to the specified Amazon S3 location. For more information, see [Uploading Objects]in the
// Amazon S3 Developer Guide. After you upload build files to the Amazon GameLift
// Servers Amazon S3 location, you can't update them.
//
// If successful, this operation creates a new build resource with a unique build
// ID and places it in INITIALIZED status. A build must be in READY status before
// you can create fleets with it.
//
// # Learn more
//
// [Uploading Your Game]
//
// [Create a Build with Files in Amazon S3]
//
// [All APIs by task]
//
// [Create a Build with Files in Amazon S3]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build
// [Uploading Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UploadingObjects.html
// [Uploading Your Game]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html
// [upload-build]: https://docs.aws.amazon.com/cli/latest/reference/gamelift/upload-build.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
