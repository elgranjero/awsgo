package gameliftstreams

// AddStreamGroupLocations is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Add locations that can host stream sessions. To add a location, the stream
//
// group must be in ACTIVE status. You configure locations and their corresponding
// capacity for each stream group. Creating a stream group in a location that's
// nearest to your end users can help minimize latency and improve quality.
//
// This operation provisions stream capacity at the specified locations. By
// default, all locations have 1 or 2 capacity, depending on the stream class
// option: 2 for 'High' and 1 for 'Ultra' and 'Win2022'. This operation also copies
// the content files of all associated applications to an internal S3 bucket at
// each location. This allows Amazon GameLift Streams to host performant stream
// sessions.
