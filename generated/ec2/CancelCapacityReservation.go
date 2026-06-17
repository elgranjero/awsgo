package ec2

// CancelCapacityReservation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Cancels the specified Capacity Reservation, releases the reserved capacity, and
// changes the Capacity Reservation's state to cancelled .
//
// You can cancel a Capacity Reservation that is in the following states:
//
// - assessing
//
// - active and there is no commitment duration or the commitment duration has
// elapsed. You can't cancel a future-dated Capacity Reservation during the
// commitment duration.
//
// You can't modify or cancel a Capacity Block. For more information, see [Capacity Blocks for ML].
//
// If a future-dated Capacity Reservation enters the delayed state, the commitment
// duration is waived, and you can cancel it as soon as it enters the active state.
//
// Instances running in the reserved capacity continue running until you stop
// them. Stopped instances that target the Capacity Reservation can no longer
// launch. Modify these instances to either target a different Capacity
// Reservation, launch On-Demand Instance capacity, or run in any open Capacity
// Reservation that has matching attributes and sufficient capacity.
//
// [Capacity Blocks for ML]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-blocks.html
