package greengrassv2

// ListCoreDevices is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Retrieves a paginated list of Greengrass core devices.
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
// - For Greengrass nucleus 2.12.2 and earlier, the core device sends status
// updates when the status of any component on the core device becomes ERRORED or
// BROKEN .
//
// - For Greengrass nucleus 2.12.3 and later, the core device sends status
// updates when the status of any component on the core device becomes ERRORED ,
// BROKEN , RUNNING , or FINISHED .
//
// - At a [regular interval that you can configure], which defaults to 24 hours
//
// - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
// local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
