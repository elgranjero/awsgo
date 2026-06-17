package directoryservice

// UpdateHybridAD is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Updates the configuration of an existing hybrid directory. You can recover
// hybrid directory administrator account or modify self-managed instance settings.
//
// Updates are applied asynchronously. Use DescribeHybridADUpdate to monitor the progress of
// configuration changes.
//
// The InstanceIds must have a one-to-one correspondence with CustomerDnsIps ,
// meaning that if the IP address for instance i-10243410 is 10.24.34.100 and the
// IP address for instance i-10243420 is 10.24.34.200, then the input arrays must
// maintain the same order relationship, either [10.24.34.100, 10.24.34.200] paired
// with [i-10243410, i-10243420] or [10.24.34.200, 10.24.34.100] paired with
// [i-10243420, i-10243410].
//
// You must provide at least one update to UpdateHybridADRequest$HybridAdministratorAccountUpdate or UpdateHybridADRequest$SelfManagedInstancesSettings.
