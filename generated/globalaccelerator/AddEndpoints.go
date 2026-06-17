package globalaccelerator

// AddEndpoints is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Add endpoints to an endpoint group. The AddEndpoints API operation is the
// recommended option for adding endpoints. The alternative options are to add
// endpoints when you create an endpoint group (with the [CreateEndpointGroup]API) or when you update
// an endpoint group (with the [UpdateEndpointGroup]API).
//
// There are two advantages to using AddEndpoints to add endpoints in Global
// Accelerator:
//
// - It's faster, because Global Accelerator only has to resolve the new
// endpoints that you're adding, rather than resolving new and existing endpoints.
//
// - It's more convenient, because you don't need to specify the current
// endpoints that are already in the endpoint group, in addition to the new
// endpoints that you want to add.
//
// For information about endpoint types and requirements for endpoints that you
// can add to Global Accelerator, see [Endpoints for standard accelerators]in the Global Accelerator Developer Guide.
//
// [Endpoints for standard accelerators]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints.html
// [UpdateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_UpdateEndpointGroup.html
// [CreateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_CreateEndpointGroup.html
