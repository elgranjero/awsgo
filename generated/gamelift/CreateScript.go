package gamelift

// CreateScript is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere
//
// Creates a new script record for your Amazon GameLift Servers Realtime script.
// Realtime scripts are JavaScript that provide configuration settings and optional
// custom game logic for your game. The script is deployed when you create a Amazon
// GameLift Servers Realtime fleet to host your game sessions. Script logic is
// executed during an active game session.
//
// To create a new script record, specify a script name and provide the script
// file(s). The script files and all dependencies must be zipped into a single
// file. You can pull the zip file from either of these locations:
//
// - A locally available directory. Use the ZipFile parameter for this option.
//
// - An Amazon Simple Storage Service (Amazon S3) bucket under your Amazon Web
// Services account. Use the StorageLocation parameter for this option. You'll need
// to have an Identity Access Management (IAM) role that allows the Amazon GameLift
// Servers service to access your S3 bucket.
//
// If the call is successful, a new script record is created with a unique script
// ID. If the script file is provided as a local file, the file is uploaded to an
// Amazon GameLift Servers-owned S3 bucket and the script record's storage location
// reflects this location. If the script file is provided as an S3 bucket, Amazon
// GameLift Servers accesses the file at this storage location as needed for
// deployment.
//
// # Learn more
//
// [Amazon GameLift Servers Amazon GameLift Servers Realtime]
//
// [Set Up a Role for Amazon GameLift Servers Access]
//
// # Related actions
//
// [All APIs by task]
//
// [Set Up a Role for Amazon GameLift Servers Access]: https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html
// [Amazon GameLift Servers Amazon GameLift Servers Realtime]: https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
