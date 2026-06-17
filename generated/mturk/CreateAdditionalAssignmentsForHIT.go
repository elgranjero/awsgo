package mturk

// CreateAdditionalAssignmentsForHIT is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The CreateAdditionalAssignmentsForHIT operation increases the maximum number
//
// of assignments of an existing HIT.
//
// To extend the maximum number of assignments, specify the number of additional
// assignments.
//
// - HITs created with fewer than 10 assignments cannot be extended to have 10
// or more assignments. Attempting to add assignments in a way that brings the
// total number of assignments for a HIT from fewer than 10 assignments to 10 or
// more assignments will result in an
// AWS.MechanicalTurk.InvalidMaximumAssignmentsIncrease exception.
//
// - HITs that were created before July 22, 2015 cannot be extended. Attempting
// to extend HITs that were created before July 22, 2015 will result in an
// AWS.MechanicalTurk.HITTooOldForExtension exception.
