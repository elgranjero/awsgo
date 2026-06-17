package gameliftstreams

// UpdateStreamGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Updates the configuration settings for an Amazon GameLift Streams stream group
//
// resource. To update a stream group, it must be in ACTIVE status. You can change
// the description, the set of locations, and the requested capacity of a stream
// group per location. If you want to change the stream class, create a new stream
// group.
//
// Stream capacity represents the number of concurrent streams that can be active
// at a time. You set stream capacity per location, per stream group. The following
// capacity settings are available:
//
// - Always-on capacity: This setting, if non-zero, indicates minimum streaming
// capacity which is allocated to you and is never released back to the service.
// You pay for this base level of capacity at all times, whether used or idle.
//
// - Maximum capacity: This indicates the maximum capacity that the service can
// allocate for you. Newly created streams may take a few minutes to start.
// Capacity is released back to the service when idle. You pay for capacity that is
// allocated to you until it is released.
//
// - Target-idle capacity: This indicates idle capacity which the service
// pre-allocates and holds for you in anticipation of future activity. This helps
// to insulate your users from capacity-allocation delays. You pay for capacity
// which is held in this intentional idle state.
//
// Values for capacity must be whole number multiples of the tenancy value of the
// stream group's stream class.
//
// To update a stream group, specify the stream group's Amazon Resource Name (ARN)
// and provide the new values. If the request is successful, Amazon GameLift
// Streams returns the complete updated metadata for the stream group. Expired
// stream groups cannot be updated.
