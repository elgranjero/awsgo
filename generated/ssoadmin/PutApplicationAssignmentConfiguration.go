package ssoadmin

// PutApplicationAssignmentConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Configure how users gain access to an application. If AssignmentsRequired is
// true (default value), users don’t have access to the application unless an
// assignment is created using the [CreateApplicationAssignment API]. If false , all users have access to the
// application. If an assignment is created using [CreateApplicationAssignment]., the user retains access if
// AssignmentsRequired is set to true .
//
// [CreateApplicationAssignment API]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_CreateApplicationAssignment.html
// [CreateApplicationAssignment]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_CreateApplicationAssignment.html
