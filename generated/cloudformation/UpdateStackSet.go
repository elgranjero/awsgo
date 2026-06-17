package cloudformation

// UpdateStackSet is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Updates the StackSet and associated stack instances in the specified accounts
// and Amazon Web Services Regions.
//
// Even if the StackSet operation created by updating the StackSet fails
// (completely or partially, below or above a specified failure tolerance), the
// StackSet is updated with your changes. Subsequent CreateStackInstancescalls on the specified
// StackSet use the updated StackSet.
//
// The maximum number of organizational unit (OUs) supported by a UpdateStackSet
// operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
