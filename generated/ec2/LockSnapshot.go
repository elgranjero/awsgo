package ec2

// LockSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Locks an Amazon EBS snapshot in either governance or compliance mode to protect
// it against accidental or malicious deletions for a specific duration. A locked
// snapshot can't be deleted.
//
// You can also use this action to modify the lock settings for a snapshot that is
// already locked. The allowed modifications depend on the lock mode and lock
// state:
//
// - If the snapshot is locked in governance mode, you can modify the lock mode
// and the lock duration or lock expiration date.
//
// - If the snapshot is locked in compliance mode and it is in the cooling-off
// period, you can modify the lock mode and the lock duration or lock expiration
// date.
//
// - If the snapshot is locked in compliance mode and the cooling-off period has
// lapsed, you can only increase the lock duration or extend the lock expiration
// date.
