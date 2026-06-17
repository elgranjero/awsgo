package ec2

// ModifyInstancePlacement is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the placement attributes for a specified instance. You can do the
// following:
//
// - Modify the affinity between an instance and a [Dedicated Host]. When affinity is set to host
// and the instance is not associated with a specific Dedicated Host, the next time
// the instance is started, it is automatically associated with the host on which
// it lands. If the instance is restarted or rebooted, this relationship persists.
//
// - Change the Dedicated Host with which an instance is associated.
//
// - Change the instance tenancy of an instance.
//
// - Move an instance to or from a [placement group].
//
// At least one attribute for affinity, host ID, tenancy, or placement group name
// must be specified in the request. Affinity and tenancy can be modified in the
// same request.
//
// To modify the host ID, tenancy, placement group, or partition for an instance,
// the instance must be in the stopped state.
//
// [Dedicated Host]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html
// [placement group]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
