package cloudformation

// DeleteStackInstances is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Deletes stack instances for the specified accounts, in the specified Amazon Web
// Services Regions.
//
// The maximum number of organizational unit (OUs) supported by a
// DeleteStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
