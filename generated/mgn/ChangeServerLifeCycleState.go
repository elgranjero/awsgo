package mgn

// ChangeServerLifeCycleState is generated as a reference stub.
// Executable command wiring lives under cmd/mgn.go.
//
// Allows the user to set the SourceServer.LifeCycle.state property for specific
// Source Server IDs to one of the following: READY_FOR_TEST or READY_FOR_CUTOVER.
// This command only works if the Source Server is already launchable
// (dataReplicationInfo.lagDuration is not null.)
