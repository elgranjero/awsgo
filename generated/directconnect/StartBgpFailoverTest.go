package directconnect

// StartBgpFailoverTest is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Starts the virtual interface failover test that verifies your configuration
// meets your resiliency requirements by placing the BGP peering session in the
// DOWN state. You can then send traffic to verify that there are no outages.
//
// You can run the test on public, private, transit, and hosted virtual interfaces.
//
// You can use [ListVirtualInterfaceTestHistory] to view the virtual interface test history.
//
// If you need to stop the test before the test interval completes, use [StopBgpFailoverTest].
//
// [ListVirtualInterfaceTestHistory]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_ListVirtualInterfaceTestHistory.html
// [StopBgpFailoverTest]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_StopBgpFailoverTest.html
