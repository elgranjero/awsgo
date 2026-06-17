package gamelift

// DeleteScript is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Deletes a Realtime script. This operation permanently deletes the script
// record. If script files were uploaded, they are also deleted (files stored in an
// S3 bucket are not deleted).
//
// To delete a script, specify the script ID. Before deleting a script, be sure to
// terminate all fleets that are deployed with the script being deleted. Fleet
// instances periodically check for script updates, and if the script record no
// longer exists, the instance will go into an error state and be unable to host
// game sessions.
//
// # Learn more
//
// [Amazon GameLift Servers Amazon GameLift Servers Realtime]
//
// # Related actions
//
// [All APIs by task]
//
// [Amazon GameLift Servers Amazon GameLift Servers Realtime]: https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
