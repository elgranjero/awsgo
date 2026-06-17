package greengrassv2

// ResolveComponentCandidates is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Retrieves a list of components that meet the component, version, and platform
// requirements of a deployment. Greengrass core devices call this operation when
// they receive a deployment to identify the components to install.
//
// This operation identifies components that meet all dependency requirements for
// a deployment. If the requirements conflict, then this operation returns an error
// and the deployment fails. For example, this occurs if component A requires
// version >2.0.0 and component B requires version <2.0.0 of a component
// dependency.
//
// When you specify the component candidates to resolve, IoT Greengrass compares
// each component's digest from the core device with the component's digest in the
// Amazon Web Services Cloud. If the digests don't match, then IoT Greengrass
// specifies to use the version from the Amazon Web Services Cloud.
//
// To use this operation, you must use the data plane API endpoint and
// authenticate with an IoT device certificate. For more information, see [IoT Greengrass endpoints and quotas].
//
// [IoT Greengrass endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/greengrass.html
