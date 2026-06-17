package cloudhsmv2

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudhsmv2.go.
//
// Creates or updates an CloudHSM resource policy. A resource policy helps you to
// define the IAM entity (for example, an Amazon Web Services account) that can
// manage your CloudHSM resources. The following resources support CloudHSM
// resource policies:
//
// - Backup - The resource policy allows you to describe the backup and restore
// a cluster from the backup in another Amazon Web Services account.
//
// In order to share a backup, it must be in a 'READY' state and you must own it.
//
// While you can share a backup using the CloudHSM PutResourcePolicy operation, we
// recommend using Resource Access Manager (RAM) instead. Using RAM provides
// multiple benefits as it creates the policy for you, allows multiple resources to
// be shared at one time, and increases the discoverability of shared resources. If
// you use PutResourcePolicy and want consumers to be able to describe the backups
// you share with them, you must promote the backup to a standard RAM Resource
// Share using the RAM PromoteResourceShareCreatedFromPolicy API operation.
//
// For more information, see [Working with shared backups] in the CloudHSM User Guide
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
//
// [Working with shared backups]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/sharing.html
