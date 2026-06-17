package ec2

// CancelCapacityReservationFleets is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Cancels one or more Capacity Reservation Fleets. When you cancel a Capacity
// Reservation Fleet, the following happens:
//
// - The Capacity Reservation Fleet's status changes to cancelled .
//
// - The individual Capacity Reservations in the Fleet are cancelled. Instances
// running in the Capacity Reservations at the time of cancelling the Fleet
// continue to run in shared capacity.
//
// - The Fleet stops creating new Capacity Reservations.
