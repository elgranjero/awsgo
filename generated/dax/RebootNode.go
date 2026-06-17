package dax

// RebootNode is generated as a reference stub.
// Executable command wiring lives under cmd/dax.go.
//
// Reboots a single node of a DAX cluster. The reboot action takes place as soon
// as possible. During the reboot, the node status is set to REBOOTING.
//
// RebootNode restarts the DAX engine process and does not remove the contents of
// the cache.
