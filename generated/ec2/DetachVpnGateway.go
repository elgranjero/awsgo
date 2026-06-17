package ec2

// DetachVpnGateway is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Detaches a virtual private gateway from a VPC. You do this if you're planning
// to turn off the VPC and not use it anymore. You can confirm a virtual private
// gateway has been completely detached from a VPC by describing the virtual
// private gateway (any attachments to the virtual private gateway are also
// described).
//
// You must wait for the attachment's state to switch to detached before you can
// delete the VPC or attach a different VPC to the virtual private gateway.
