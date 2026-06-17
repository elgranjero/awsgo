package storagegateway

// UpdateGatewaySoftwareNow is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update.
//
// When you make this request, you get a 200 OK success response immediately.
// However, it might take some time for the update to complete. You can call DescribeGatewayInformationto
// verify the gateway is in the STATE_RUNNING state.
//
// A software update forces a system restart of your gateway. You can minimize the
// chance of any disruption to your applications by increasing your iSCSI
// Initiators' timeouts. For more information about increasing iSCSI Initiator
// timeouts for Windows and Linux, see [Customizing your Windows iSCSI settings]and [Customizing your Linux iSCSI settings], respectively.
//
// [Customizing your Linux iSCSI settings]: https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorRedHatClient.html#CustomizeLinuxiSCSISettings
// [Customizing your Windows iSCSI settings]: https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorWindowsClient.html#CustomizeWindowsiSCSISettings
