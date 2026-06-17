package directconnect

// UpdateLag is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Updates the attributes of the specified link aggregation group (LAG).
//
// You can update the following LAG attributes:
//
// - The name of the LAG.
//
// - The value for the minimum number of connections that must be operational
// for the LAG itself to be operational.
//
// - The LAG's MACsec encryption mode.
//
// Amazon Web Services assigns this value to each connection which is part of the
//
// LAG.
//
// - The tags
//
// If you adjust the threshold value for the minimum number of operational
// connections, ensure that the new value does not cause the LAG to fall below the
// threshold and become non-operational.
