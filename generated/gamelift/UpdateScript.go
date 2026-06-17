package gamelift

// UpdateScript is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Updates Realtime script metadata and content.
//
// To update script metadata, specify the script ID and provide updated name
// and/or version values.
//
// To update script content, provide an updated zip file by pointing to either a
// local file or an Amazon S3 bucket location. You can use either method regardless
// of how the original script was uploaded. Use the Version parameter to track
// updates to the script.
//
// If the call is successful, the updated metadata is stored in the script record
// and a revised script is uploaded to the Amazon GameLift Servers service. Once
// the script is updated and acquired by a fleet instance, the new version is used
// for all new game sessions.
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
