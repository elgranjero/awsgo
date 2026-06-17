package greengrassv2

// ListInstalledComponents is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Retrieves a paginated list of the components that a Greengrass core device
// runs. By default, this list doesn't include components that are deployed as
// dependencies of other components. To include dependencies in the response, set
// the topologyFilter parameter to ALL .
//
// IoT Greengrass relies on individual devices to send status updates to the
// Amazon Web Services Cloud. If the IoT Greengrass Core software isn't running on
// the device, or if device isn't connected to the Amazon Web Services Cloud, then
// the reported status of that device might not reflect its current status. The
// status timestamp indicates when the device status was last updated.
//
// Core devices send status updates at the following times:
//
// - When the IoT Greengrass Core software starts
//
// - When the core device receives a deployment from the Amazon Web Services
// Cloud
//
// - When the status of any component on the core device becomes BROKEN
//
// - At a [regular interval that you can configure], which defaults to 24 hours
//
// - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
// local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
