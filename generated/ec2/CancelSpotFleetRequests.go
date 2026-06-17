package ec2

// CancelSpotFleetRequests is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Cancels the specified Spot Fleet requests.
//
// After you cancel a Spot Fleet request, the Spot Fleet launches no new instances.
//
// You must also specify whether a canceled Spot Fleet request should terminate
// its instances. If you choose to terminate the instances, the Spot Fleet request
// enters the cancelled_terminating state. Otherwise, the Spot Fleet request
// enters the cancelled_running state and the instances continue to run until they
// are interrupted or you terminate them manually.
//
// Terminating an instance is permanent and irreversible.
//
// After you terminate an instance, you can no longer connect to it, and it can't
// be recovered. All attached Amazon EBS volumes that are configured to be deleted
// on termination are also permanently deleted and can't be recovered. All data
// stored on instance store volumes is permanently lost. For more information, see [How instance termination works]
// .
//
// Before you terminate an instance, ensure that you have backed up all data that
// you need to retain after the termination to persistent storage.
//
// Restrictions
//
// - You can delete up to 100 fleets in a single request. If you exceed the
// specified number, no fleets are deleted.
//
// [How instance termination works]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-ec2-instance-termination-works.html
