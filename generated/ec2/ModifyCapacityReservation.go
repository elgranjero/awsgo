package ec2

// ModifyCapacityReservation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies a Capacity Reservation's capacity, instance eligibility, and the
// conditions under which it is to be released. You can't modify a Capacity
// Reservation's instance type, EBS optimization, platform, instance store
// settings, Availability Zone, or tenancy. If you need to modify any of these
// attributes, we recommend that you cancel the Capacity Reservation, and then
// create a new one with the required attributes. For more information, see [Modify an active Capacity Reservation].
//
// The allowed modifications depend on the state of the Capacity Reservation:
//
// - assessing or scheduled state - You can modify the tags only.
//
// - pending state - You can't modify the Capacity Reservation in any way.
//
// - active state but still within the commitment duration - You can't decrease
// the instance count or set an end date that is within the commitment duration.
// All other modifications are allowed.
//
// - active state with no commitment duration or elapsed commitment duration -
// All modifications are allowed.
//
// - expired , cancelled , unsupported , or failed state - You can't modify the
// Capacity Reservation in any way.
//
// [Modify an active Capacity Reservation]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-modify.html
