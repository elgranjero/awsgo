package ec2

// DeleteFleets is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes the specified EC2 Fleet request.
//
// After you delete an EC2 Fleet request, it launches no new instances.
//
// You must also specify whether a deleted EC2 Fleet request should terminate its
// instances. If you choose to terminate the instances, the EC2 Fleet request
// enters the deleted_terminating state. Otherwise, it enters the deleted_running
// state, and the instances continue to run until they are interrupted or you
// terminate them manually.
//
// A deleted instant fleet with running instances is not supported. When you
// delete an instant fleet, Amazon EC2 automatically terminates all its instances.
// For fleets with more than 1000 instances, the deletion request might fail. If
// your fleet has more than 1000 instances, first terminate most of the instances
// manually, leaving 1000 or fewer. Then delete the fleet, and the remaining
// instances will be terminated automatically.
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
// - You can delete up to 25 fleets of type instant in a single request.
//
// - You can delete up to 100 fleets of type maintain or request in a single
// request.
//
// - You can delete up to 125 fleets in a single request, provided you do not
// exceed the quota for each fleet type, as specified above.
//
// - If you exceed the specified number of fleets to delete, no fleets are
// deleted.
//
// For more information, see [Delete an EC2 Fleet request and the instances in the fleet] in the Amazon EC2 User Guide.
//
// [How instance termination works]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-ec2-instance-termination-works.html
// [Delete an EC2 Fleet request and the instances in the fleet]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/delete-fleet.html
