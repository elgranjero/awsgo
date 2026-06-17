package gamelift

// GetGameSessionLogUrl is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves the location of stored game session logs for a specified game session
// on Amazon GameLift Servers managed fleets. When a game session is terminated,
// Amazon GameLift Servers automatically stores the logs in Amazon S3 and retains
// them for 14 days. Use this URL to download the logs.
//
// See the [Amazon Web Services Service Limits] page for maximum log file sizes. Log files that exceed this limit are
// not saved.
//
// [All APIs by task]
//
// [Amazon Web Services Service Limits]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_gamelift
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
