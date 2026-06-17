package ecs

// DeleteService is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deletes a specified service within a cluster. You can delete a service if you
// have no running tasks in it and the desired task count is zero. If the service
// is actively maintaining tasks, you can't delete it, and you must update the
// service to a desired task count of zero. For more information, see [UpdateService].
//
// When you delete a service, if there are still running tasks that require
// cleanup, the service status moves from ACTIVE to DRAINING , and the service is
// no longer visible in the console or in the [ListServices]API operation. After all tasks have
// transitioned to either STOPPING or STOPPED status, the service status moves
// from DRAINING to INACTIVE . Services in the DRAINING or INACTIVE status can
// still be viewed with the [DescribeServices]API operation. However, in the future, INACTIVE
// services may be cleaned up and purged from Amazon ECS record keeping, and [DescribeServices]calls
// on those services return a ServiceNotFoundException error.
//
// If you attempt to create a new service with the same name as an existing
// service in either ACTIVE or DRAINING status, you receive an error.
//
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [ListServices]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListServices.html
// [DescribeServices]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeServices.html
