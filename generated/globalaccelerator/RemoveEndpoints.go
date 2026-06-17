package globalaccelerator

// RemoveEndpoints is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Remove endpoints from an endpoint group.
//
// The RemoveEndpoints API operation is the recommended option for removing
// endpoints. The alternative is to remove endpoints by updating an endpoint group
// by using the [UpdateEndpointGroup]API operation. There are two advantages to using AddEndpoints to
// remove endpoints instead:
//
// - It's more convenient, because you only need to specify the endpoints that
// you want to remove. With the UpdateEndpointGroup API operation, you must
// specify all of the endpoints in the endpoint group except the ones that you want
// to remove from the group.
//
// - It's faster, because Global Accelerator doesn't need to resolve any
// endpoints. With the UpdateEndpointGroup API operation, Global Accelerator must
// resolve all of the endpoints that remain in the group.
//
// [UpdateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_UpdateEndpointGroup.html
