package ssoadmin

// DescribeInstance is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Returns the details of an instance of IAM Identity Center. The status can be
// one of the following:
//
// - CREATE_IN_PROGRESS - The instance is in the process of being created. When
// the instance is ready for use, DescribeInstance returns the status of ACTIVE .
// While the instance is in the CREATE_IN_PROGRESS state, you can call only
// DescribeInstance and DeleteInstance operations.
//
// - DELETE_IN_PROGRESS - The instance is being deleted. Returns
// AccessDeniedException after the delete operation completes.
//
// - ACTIVE - The instance is active.
