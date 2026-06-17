package storagegateway

// ShutdownGateway is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Shuts down a Tape Gateway or Volume Gateway. To specify which gateway to shut
// down, use the Amazon Resource Name (ARN) of the gateway in the body of your
// request.
//
// This API action cannot be used to shut down S3 File Gateway or FSx File Gateway.
//
// The operation shuts down the gateway service component running in the gateway's
// virtual machine (VM) and not the host VM.
//
// If you want to shut down the VM, it is recommended that you first shut down the
// gateway component in the VM to avoid unpredictable conditions.
//
// After the gateway is shutdown, you cannot call any other API except StartGateway, DescribeGatewayInformation, and ListGateways.
// For more information, see ActivateGateway. Your applications cannot read from or write to the
// gateway's storage volumes, and there are no snapshots taken.
//
// When you make a shutdown request, you will get a 200 OK success response
// immediately. However, it might take some time for the gateway to shut down. You
// can call the DescribeGatewayInformationAPI to check the status. For more information, see ActivateGateway.
//
// If do not intend to use the gateway again, you must delete the gateway (using DeleteGateway)
// to no longer pay software charges associated with the gateway.
