package ssm

// CreateAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// A State Manager association defines the state that you want to maintain on your
// managed nodes. For example, an association can specify that anti-virus software
// must be installed and running on your managed nodes, or that certain ports must
// be closed. For static targets, the association specifies a schedule for when the
// configuration is reapplied. For dynamic targets, such as an Amazon Web Services
// resource group or an Amazon Web Services autoscaling group, State Manager, a
// tool in Amazon Web Services Systems Manager applies the configuration when new
// managed nodes are added to the group. The association also specifies actions to
// take when applying the configuration. For example, an association for anti-virus
// software might run once a day. If the software isn't installed, then State
// Manager installs it. If the software is installed, but the service isn't
// running, then the association might instruct State Manager to start the service.
