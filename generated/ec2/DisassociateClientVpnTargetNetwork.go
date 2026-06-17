package ec2

// DisassociateClientVpnTargetNetwork is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates a target network from the specified Client VPN endpoint. When you
// disassociate the last target network from a Client VPN, the following happens:
//
// - The route that was automatically added for the VPC is deleted
//
// - All active client connections are terminated
//
// - New client connections are disallowed
//
// - The Client VPN endpoint's status changes to pending-associate
