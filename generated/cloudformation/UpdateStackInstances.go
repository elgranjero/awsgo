package cloudformation

// UpdateStackInstances is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Updates the parameter values for stack instances for the specified accounts,
// within the specified Amazon Web Services Regions. A stack instance refers to a
// stack in a specific account and Region.
//
// You can only update stack instances in Amazon Web Services Regions and accounts
// where they already exist; to create additional stack instances, use [CreateStackInstances].
//
// During StackSet updates, any parameters overridden for a stack instance aren't
// updated, but retain their overridden value.
//
// You can only update the parameter values that are specified in the StackSet. To
// add or delete a parameter itself, use [UpdateStackSet]to update the StackSet template. If you
// add a parameter to a template, before you can override the parameter value
// specified in the StackSet you must first use [UpdateStackSet]to update all stack instances with
// the updated template and parameter value specified in the StackSet. Once a stack
// instance has been updated with the new parameter, you can then override the
// parameter value using UpdateStackInstances .
//
// The maximum number of organizational unit (OUs) supported by a
// UpdateStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
//
// [CreateStackInstances]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html
// [UpdateStackSet]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html
