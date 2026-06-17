package gameliftstreams

// CreateStreamGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Stream groups manage how Amazon GameLift Streams allocates resources and
//
// handles concurrent streams, allowing you to effectively manage capacity and
// costs. Within a stream group, you specify an application to stream, streaming
// locations and their capacity, and the stream class you want to use when
// streaming applications to your end-users. A stream class defines the hardware
// configuration of the compute resources that Amazon GameLift Streams will use
// when streaming, such as the CPU, GPU, and memory.
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
// To adjust the capacity of any ACTIVE stream group, call [UpdateStreamGroup].
//
// If the CreateStreamGroup request is successful, Amazon GameLift Streams assigns
// a unique ID to the stream group resource and sets the status to ACTIVATING . It
// can take a few minutes for Amazon GameLift Streams to finish creating the stream
// group while it searches for unallocated compute resources and provisions them.
// When complete, the stream group status will be ACTIVE and you can start stream
// sessions by using [StartStreamSession]. To check the stream group's status, call [GetStreamGroup].
//
// Stream groups should be recreated every 3-4 weeks to pick up important service
// updates and fixes. Stream groups that are older than 180 days can no longer be
// updated with new application associations. Stream groups expire when they are
// 365 days old, at which point they can no longer stream sessions. The exact
// expiration date is indicated by the date value in the ExpiresAt field.
//
// [GetStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamGroup.html
// [UpdateStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_UpdateStreamGroup.html
// [StartStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_StartStreamSession.html
