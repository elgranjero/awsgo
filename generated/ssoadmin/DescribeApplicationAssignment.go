package ssoadmin

// DescribeApplicationAssignment is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Retrieves a direct assignment of a user or group to an application. If the user
// doesn’t have a direct assignment to the application, the user may still have
// access to the application through a group. Therefore, don’t use this API to test
// access to an application for a user. Instead use ListApplicationAssignmentsForPrincipal.
