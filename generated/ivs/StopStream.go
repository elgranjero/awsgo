package ivs

// StopStream is generated as a reference stub.
// Executable command wiring lives under cmd/ivs.go.
//
// Disconnects the incoming RTMPS stream for the specified channel. Can be used in
// conjunction with DeleteStreamKeyto prevent further streaming to a channel.
//
// Many streaming client-software libraries automatically reconnect a dropped
// RTMPS session, so to stop the stream permanently, you may want to first revoke
// the streamKey attached to the channel.
