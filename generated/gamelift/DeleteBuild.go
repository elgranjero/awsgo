package gamelift

// DeleteBuild is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Deletes a build. This operation permanently deletes the build resource and any
// uploaded build files. Deleting a build does not affect the status of any active
// fleets using the build, but you can no longer create new fleets with the deleted
// build.
//
// To delete a build, specify the build ID.
//
// # Learn more
//
// [Upload a Custom Server Build]
//
// [All APIs by task]
//
// [Upload a Custom Server Build]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
