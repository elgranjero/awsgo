package gameliftstreams

// RemoveStreamGroupLocations is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Removes a set of remote locations from this stream group. To remove a
//
// location, the stream group must be in ACTIVE status. When you remove a
// location, Amazon GameLift Streams releases allocated compute resources in that
// location. Stream sessions can no longer start from removed locations in a stream
// group. Amazon GameLift Streams also deletes the content files of all associated
// applications that were in Amazon GameLift Streams's internal Amazon S3 bucket at
// this location.
//
// You cannot remove the Amazon Web Services Region location where you initially
// created this stream group, known as the primary location. However, you can set
// the stream capacity to zero to avoid incurring costs for allocated compute
// resources in that location.
