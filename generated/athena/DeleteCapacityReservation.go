package athena

// DeleteCapacityReservation is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Deletes a cancelled capacity reservation. A reservation must be cancelled
// before it can be deleted. A deleted reservation is immediately removed from your
// account and can no longer be referenced, including by its ARN. A deleted
// reservation cannot be called by GetCapacityReservation , and deleted
// reservations do not appear in the output of ListCapacityReservations .
