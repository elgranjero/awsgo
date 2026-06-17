package cloudformation

// CreateStackInstances is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Creates stack instances for the specified accounts, within the specified Amazon
// Web Services Regions. A stack instance refers to a stack in a specific account
// and Region. You must specify at least one value for either Accounts or
// DeploymentTargets , and you must specify at least one value for Regions .
//
// The maximum number of organizational unit (OUs) supported by a
// CreateStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
